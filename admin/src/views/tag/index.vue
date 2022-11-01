<template>
  <el-container>
    <el-header>
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-button type="danger" icon="el-icon-delete" size="small" style="margin-top: 20px " @click="handleDeleteList">批量删除</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="el-icon-plus" size="small" style="margin-top: 20px" @click="addFormVisible = true">新增</el-button>
        </el-form-item>
        <el-form-item>
          <el-input v-model="key" size="small" placeholder="请输入名称" style="margin-top: 16px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="el-icon-search" size="small" style="margin-top: 20px" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" style="margin-top: 20px" @click="reset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-header>
    <el-main><el-table
      ref="multipleTable"
      :data="pageList"
      tooltip-effect="dark"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        sortable="custom"
        align="center"
        type="selection"
        width="55"
      />
      <el-table-column
        prop="Id"
        label="Id"
        align="center"
        width="120"
      />
      <el-table-column
        prop="Name"
        label="标签名"
        align="center"
        width="160"
      />
      <el-table-column
        prop="Description"
        align="center"
        label="描述 "
        show-overflow-tooltip
      />
      <el-table-column
        label="创建时间"
        align="center"
        width="189"
      >
        <template slot-scope="scope">{{ scope.row.CreatedAt }}</template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <!--          <el-button-->
          <!--            size="mini"-->
          <!--            @click="handleEdit(scope.$index, scope.row)">编辑-->
          <!--          </el-button>-->
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table></el-main>
    <el-footer>
      <div class="block">
        <el-pagination
          :current-page.sync="currentPage"
          :page-size="pageSize"
          layout="prev, pager, next, jumper"
          :total="categoryCount"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-footer>
    <!--    修改标签-->
    <el-dialog title="修改标签" :visible.sync="editFormVisible">
      <el-form :model="editForm">
        <el-form-item label="标签名" :label-width="formLabelWidth">
          <el-input v-model="editForm.Name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="标签描述" :label-width="formLabelWidth">
          <el-input v-model="editForm.Description" autocomplete="off" />
        </el-form-item>
        <!--        <el-form-item label="父级分类" :label-width="formLabelWidth">-->
        <!--          <el-select v-model="editForm.ParentId" placeholder="请选择父级分类">-->
        <!--            <el-option label="区域一" value="shanghai"></el-option>-->
        <!--            <el-option label="区域二" value="beijing"></el-option>-->
        <!--          </el-select>-->
        <!--        </el-form-item>-->
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="editFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleConfirmEdit">确 定</el-button>
      </div>
    </el-dialog>
    <!--    添加标签-->
    <el-dialog title="添加标签" :visible.sync="addFormVisible">
      <el-form :model="addForm">
        <el-form-item label="标签名" :label-width="formLabelWidth">
          <el-input v-model="addForm.Name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="标签描述" :label-width="formLabelWidth">
          <el-input v-model="addForm.Description" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleConfirmAdd">确 定</el-button>
      </div>
    </el-dialog>
  </el-container>
</template>

<script>
import {
  fetchTagCount,
  fetchTagList,
  updateTag,
  deleteTag,
  addTag,
  multiDelTags
} from '@/api/tag'
export default {
  name: 'Tag',
  data() {
    return {
      key: '', // 搜索关键词
      categoryCount: 2,
      currentPage: 1,
      pageSize: 10,
      pageList: [],
      tableData: [],
      editFormVisible: false,
      addFormVisible: false,
      editForm: {
      },
      addForm: {
      },
      formLabelWidth: '120px',
      selection: [] // 选中条目
    }
  },
  mounted() {
    // this.category  Count = this.tableData.length
    this.getCount()
    this.currentChangePage(1)
  },
  methods: {
    getCount() {
      fetchTagCount().then(response => {
        const { data } = response
        this.categoryCount = data
      })
    },
    // 重置
    reset() {
      this.currentPage = 1 // 页码
      this.pageSize = 10 // 每页条数
      this.getCount() // 总条数
      this.currentChangePage(1)
      this.key = '' // 搜索关键词
    },
    toggleSelection(rows) {
      this.selection = rows
      if (rows) {
        rows.forEach(row => {
          this.$refs.multipleTable.toggleRowSelection(row)
        })
      } else {
        this.$refs.multipleTable.clearSelection()
      }
    },
    handleSelectionChange(val) {
      this.selection = val
    },
    // 编辑分类信息
    handleEdit(index, row) {
      this.editFormVisible = true
      this.editForm = row
    },
    // 确定编辑分类信息 并且提交
    handleConfirmEdit(index, row) {
      this.editFormVisible = false
      var data = { 'id': this.editForm.Id, 'name': this.editForm.Name, 'description': this.editForm.Description }
      updateTag(data).then(response => {
        if (response.code === 100) {
          this.$message.success(response.msg)
        }
        this.pageList[index] = this.editForm
      })
    },
    // 添加分类确定
    handleConfirmAdd(index, row) {
      this.addFormVisible = false
      var data = { 'name': this.addForm.Name, 'description': this.addForm.Description }
      addTag(data).then(response => {
        if (response.code === 100) {
          this.$message.success(response.msg)
          this.getCount()
          this.currentChangePage(this.currentPage)
          this.addForm = {}
        }
      })
    },
    handleDelete(index, row) {
      this.$confirm('确定要删除吗?', '删除', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'error'
      }).then(() => {
        deleteTag(row.Id).then(response => {
          if (response.code === 100) {
            this.$message.success(response.msg)
            this.pageList.splice(index, 1)
            if (this.pageList.length === 0) {
              this.getCount()
              if (this.currentPage > 1) {
                this.currentChangePage(this.currentPage - 1)
              } else {
                this.currentChangePage(1)
              }
            } else {
              this.getCount()
              this.currentChangePage(this.currentPage - 1)
            }
          }
        })
      }).catch(() => {
      })
    },
    handleSizeChange: function(pageSize) { // 每页条数切换
      this.pageSize = pageSize
      this.handleCurrentChange(this.currentPage1)
    },

    handleCurrentChange: function(currentPage) { // 页码切换
      this.currentPage1 = currentPage
      this.currentChangePage(currentPage)
    },
    // 分页方法（重点）
    currentChangePage(currentPage) {
      var param = { 'page': currentPage, 'pageSize': this.pageSize }
      fetchTagList(param).then(response => {
        const { data } = response
        this.pageList = data
        for (var i = 0; i < this.pageList.length; i++) {
          var date = new Date(Date.parse(this.pageList[i].CreatedAt))
          this.pageList[i].CreatedAt = date.toLocaleString('chinese', { hour12: false })
        }
      })
    },
    // 批量删除
    handleDeleteList() {
      if (this.selection.length === 0) {
        this.$message.error('请勾选要删除的条目')
      } else {
        this.$confirm('确定要删除吗?', '删除', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error'
        }).then(() => {
          let ids = []
          this.selection.forEach(function(val) {
            ids.push(val.Id)
          })
          ids = ids.join() // 将数组转换成字符串：[1, 2, 3] --> '1, 2, 3'
          multiDelTags(ids)
            .then(res => {
              this.$message.success(res.msg)
              if (this.pageList.length === 0) {
                this.getCount()
                if (this.currentPage > 1) {
                  this.currentChangePage(this.currentPage - 1)
                } else {
                  this.currentChangePage(1)
                }
              } else {
                this.getCount()
                this.currentChangePage(this.currentPage - 1)
              }
            })
            .catch(() => {
            })
        }).catch(() => {
        })
      }
    }

  }
}
</script>

<style scoped>

</style>
