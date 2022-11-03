<template>
  <el-container>
    <el-header>
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-select v-model="pagination.category_id" style="width: 150px; margin-top: 20px" size="small" clearable placeholder="请选择分类">
            <el-option
              v-for="item in categories"
              :key="item.Id"
              :label="item.Name"
              :value="item.Id"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="pagination.state" style="width: 150px; margin-top: 20px" size="small" clearable placeholder="请选择文章状态">
            <el-option
              v-for="item in states"
              :key="item.value"
              :label="item.name"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input v-model="key" size="small" placeholder="请输入关键词" style="margin-top: 20px" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" style="margin-top: 20px" @click="search"><i class="el-icon-search" /> 搜索</el-button>
        </el-form-item>
        <el-form-item>
          <el-button size="small" style="margin-top: 20px" @click="reset"><i class="el-icon-refresh" /> 重置</el-button>
        </el-form-item>
      </el-form>
    </el-header>
    <el-main>
      <el-table
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
          prop="ID"
          label="ID"
          align="center"
          width="80"
        >
          <template slot-scope="scope">
            {{ scope.row.Post.ID }}
          </template>
        </el-table-column>
        <el-table-column
          prop="Title"
          label="文章标题"
          align="center"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            {{ scope.row.Post.Title }}
          </template>
        </el-table-column>
        <el-table-column
          prop="CategoryName"
          align="center"
          label="类别 "
          width="150"
        >
          <template slot-scope="scope">
            <el-tag type="success" effect="dark">
              {{ scope.row.Post.CategoryName }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="tagList"
          align="center"
          label="标签"
        >
          <template slot-scope="scope">
            <el-tag v-for="tag in scope.row.TagList " :key="tag.Id" style="margin-left: 5px" type="success">
              {{ tag.Name }}
            </el-tag>
          </template>

        </el-table-column>
        <el-table-column
          prop="State"
          align="center"
          label="状态 "
          width="80"
        >
          <template slot-scope="scope">
            <el-tag type="danger">
              {{ scope.row.Post.State }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          align="center"
          width="189"
        >
          <template slot-scope="scope">{{ scope.row.Post.CreatedAt }}</template>
        </el-table-column>
        <el-table-column label="操作" width="70" align="center">
          <template slot-scope="scope">
            <el-tooltip content="编辑" placement="top-start">
              <el-button
                size="small"
                icon="el-icon-edit"
                @click="handleEdit(scope.$index, scope.row)"
              />
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column label="是否发布" width="100" align="center">
          <template slot-scope="scope">
            <el-switch
              v-model="scope.row.Post.IsPublished"
              :active-value="1"
              :inactive-value="0"
              @change="handlePublish(scope.$index, scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column label="是否删除" width="100" align="center">
          <template slot-scope="scope">
            <el-switch
              v-model="scope.row.Post.IsDelete"
              :active-value="1"
              :inactive-value="0"
              @change="handleDelete(scope.$index, scope.row)"
            />
          </template>
        </el-table-column>
      </el-table></el-main>
    <el-footer>
      <div class="block">
        <el-pagination
          :current-page.sync="currentPage"
          :page-size="pageSize"
          layout="prev, pager, next, jumper"
          :total="postCount"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-footer>
    <!-- 修改文章弹窗 -->
    <el-dialog
      title="修改文章"
      :visible.sync="dialogOptions.editVisible"
      :fullscreen="dialogOptions.editFullScreen"
      width="80%"
    >
      <el-form ref="editForm" :model="editForm" :rules="editRules" label-width="80px">
        <el-form-item label="标题" prop="title">
          <el-input v-model="editForm.title" size="small" type="text" />
        </el-form-item>
        <el-row>
          <el-col :span="8">
            <el-form-item label="分类" prop="category_id">
              <!--            <el-select size="small" v-model="addForm.category_id" clearable @clear="editForm.category_id=null"-->
              <el-select
                v-model="editForm.category_id"
                size="small"
                placeholder="请选择分类"
              >
                <el-option
                  v-for="item in categories"
                  :key="item.Id"
                  :label="item.Name"
                  :value="item.Id"
                />
              </el-select>
              &nbsp;<el-button size="mini" type="primary" @click="handleOpenAddCategoryDraw">新增</el-button>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="标签" prop="selectTagIds">
              <el-select
                v-model="editForm.selectTagIds"
                size="small"
                filterable
                allow-create
                multiple
                :multiple-limit="3"
                clearable
                placeholder="请选择标签"
                @change="selectTrigger('editForm')"
              >
                <el-option
                  v-for="item in tags"
                  :key="item.Id"
                  :label="item.Name"
                  :value="item.Id"
                />
              </el-select>
              &nbsp;<el-button size="mini" type="primary" @click="handleOpenAddTagDraw">新增</el-button>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="关键词" prop="keywords">
              <el-input v-model="editForm.keywords" size="small" type="text" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="摘要" prop="description">
          <el-input v-model="editForm.description" :rows="5" type="textarea" />
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <markdown-editor ref="markdownEditor" v-model="editForm.md_content" :options="{hideModeSwitch:true,previewStyle:'tab'}" height="200px" />
        </el-form-item>
        <el-form-item>
          <el-button
            type="success"
            :loading="dialogOptions.editDraftBtnLoading"
            @click="handleRowEdit()"
          >保存
          </el-button>
          <el-button @click="dialogOptions.editVisible=false">取消</el-button>
          <el-button style="margin-top:80px;" type="primary" icon="el-icon-document" @click="getHtml">
            内容预览
          </el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-dialog
      title="文章详情"
      :visible.sync="dialogOptions.displayHTML"
      :fullscreen="false"
      :with-header="false"
      width="80%"
    >
      <!--        <span slot="footer" class="dialog-footer">-->
      <!--    <el-button @click="dialogOptions.displayHTML = false">取 消</el-button>-->
      <!--  </span>-->
      <div style="margin: 20px" v-html="html" />
    </el-dialog>
    <!-- 添加分类抽屉 -->
    <el-drawer
      title="添加分类"
      :visible.sync="drawOptions.addVisible"
      direction="rtl"
    >
      <el-form ref="addCategoryForm" :model="addCategoryForm" :rules="addCategoryRules" label-width="80px">
        <el-form-item label="分类名称" prop="name">
          <el-input v-model="addCategoryForm.name" style="width: 280px" size="small" type="text" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="addCategoryForm.description" style="width: 280px" size="small" type="text" />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            :loading="drawOptions.addBtnLoading"
            @click="handleCategoryAdd"
          >保存
          </el-button>
          <el-button @click="drawOptions.addVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 添加标签抽屉 -->
    <el-drawer
      title="添加标签"
      :visible.sync="drawOptions.addTagVisible"
      direction="rtl"
    >
      <el-form ref="addTagForm" :model="addTagForm" :rules="addTagForm" label-width="80px">
        <el-form-item label="标签名称" prop="name">
          <el-input v-model="addTagForm.name" style="width: 280px" size="small" type="text" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="addTagForm.description" style="width: 280px" size="small" type="text" />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            :loading="drawOptions.addTagBtnLoading"
            @click="handleTagAdd"
          >保存
          </el-button>
          <el-button @click="drawOptions.addTagVisible=false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </el-container>
</template>

<script>
import { addCategory, fetchAllCategory } from '@/api/category'
import {
  fetchPostCount,
  fetchPostList,
  deletePost,
  undeletePost,
  publishPost,
  stopPublishPost,
  fetchPostById,
  updatePost
} from '@/api/post'
import MarkdownEditor from '@/components/MarkdownEditor'
import { addTag, fetchAllTag } from '@/api/tag'
export default {
  name: 'Post',
  components: { MarkdownEditor },
  data() {
    return {
      tagArray: [...Array(3).keys()],
      // tagList: ['tag0','tag1','tag2'],
      key: '', // 搜索关键词
      postCount: 2,
      currentPage: 1,
      pageSize: 5,
      pageList: [],
      tableData: [],
      pagination: {
        state: null,
        category_id: null
      },
      editFormVisible: false,
      addFormVisible: false,
      editForm: {
        post_id: 0,
        author_id: null,
        category_id: null,
        selectTagIds: [],
        tag_ids: '',
        is_published: 1,
        title: '',
        description: '',
        content: '',
        md_content: '',
        keywords: ''
      },
      addForm: {
      },
      editRules: {
        user_id: [
          { required: true, trigger: 'blur', message: '用户 ID 不能为空' }
        ],
        title: [
          { required: true, trigger: 'blur', message: '请输入文章标题' },
          { max: 255, trigger: 'blur', message: 'URL 长度不能超过 255 位' }
        ],
        pwd: [
          { max: 64, trigger: 'blur', message: '密码长度不能超过 64 位' }
        ],
        url: [
          { max: 255, trigger: 'blur', message: 'URL 长度不能超过 255 位' }
        ],
        content: [
          { required: true, pattern: /^(?!\n$)/, trigger: 'blur', message: '请输入文章内容' },
          { max: 100000, trigger: 'blur', message: '内容字数不能超过 100000' }
        ],
        description: [
          { max: 255, trigger: 'blur', message: '摘要字数不能超过 255' }
        ],
        md_content: [
          { required: true, trigger: 'blur', message: '请输入文章内容' }
        ]
      },
      formLabelWidth: '120px',
      selection: [], // 选中条目
      categories: [],
      states: [
        { name: '已发布', value: 1 },
        { name: '草稿', value: 2 }
      ],
      dialogOptions: {
        displayHTML: false,
        editBtnLoading: false,
        editDraftBtnLoading: false,
        editVisible: false,
        editFullScreen: true
      },
      tags: [],
      html: '',
      drawOptions: {
        addVisible: false,
        addBtnLoading: false,
        addTagVisible: false,
        addTagBtnLoading: false
      },
      addCategoryForm: {
        name: '',
        description: ''
      },
      addCategoryRules: {
        name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
        description: [{ required: true, message: '请输入分类描述', trigger: 'blur' }]
      },
      addTagForm: {
        name: '',
        description: ''
      },
      addTagRules: {
        name: [{ required: true, message: '请输入标签名称', trigger: 'blur' }]
      }
    }
  },
  mounted() {
    this.currentChangePage(1)
    this.getCount()
    this.fetchCategoryData()
  },
  methods: {
    // 获取分类数据
    fetchCategoryData() {
      fetchAllCategory().then(response => {
        const { data } = response
        this.categories = data
      })
    },
    // 获取标签数据
    fetchTagData() {
      fetchAllTag().then(response => {
        const { data } = response
        this.tags = data
      })
    },
    getCount() {
      fetchPostCount().then(response => {
        const { data } = response
        this.postCount = data
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
    // 搜索
    search() {
      this.pagination.currentPage = 1
      // this.fetchPageData()
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
    // 清空表单验证
    clearValidate(formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].clearValidate()
      }
    },
    // 编辑文章
    handleEdit(index, row) {
      this.fetchCategoryData()
      this.fetchTagData()
      // 显示弹窗
      this.dialogOptions.editVisible = true
      // 清空表单提示信息
      this.clearValidate('editForm')
      // 初始化表单数据
      var param = { 'post_id': this.pageList[index].Post.ID }
      fetchPostById(param)
        .then(res => {
          this.editForm.post_id = this.pageList[index].Post.ID
          this.editForm.author_id = res.data.Post.AuthorId
          this.editForm.category_id = res.data.Post.CategoryId
          this.editForm.is_published = res.data.Post.IsPublished
          this.editForm.title = res.data.Post.Title
          this.editForm.description = res.data.Post.Description
          this.editForm.content = res.data.Post.Content
          this.editForm.md_content = res.data.Post.MDContent
          this.editForm.keywords = res.data.Post.Keywords
          const tagIds = []
          res.data.TagList.forEach((tag) => {
            tagIds.push(tag.Id)
          })
          this.$set(this.editForm, 'selectTagIds', tagIds)
          this.$refs.editEditor.setContent(this.editForm.content)
        })
        .catch(() => {
        })
    },
    handleDelete(index, row) {
      var param = { 'post_id': this.pageList[index].Post.ID }
      // 点按钮就已经改变状态了 现在已经是改变之后的状态
      if (this.pageList[index].Post.IsDelete === 0) {
        undeletePost(param).then(response => {
          this.$message.success(response.msg)
        })
        this.pageList[index].Post.IsDelete = 0
      } else {
        deletePost(param).then(response => {
          this.$message.success(response.msg)
        })
        this.pageList[index].Post.IsDelete = 1
      }
      // this.getCount()
      // this.currentChangePage(this.currentPage)
    },
    handlePublish(index, row) {
      // 发布的时候要判断是不是删除状态，如果是删除状态则无法发布
      if (this.pageList[index].Post.IsDelete === 1) {
        this.$message.error('文章已删除，无法发布')
        this.pageList[index].Post.IsPublished = 0
      } else {
        var param = { 'post_id': this.pageList[index].Post.ID }
        if (this.pageList[index].Post.IsPublished === 0) {
          stopPublishPost(param).then(response => {
            this.$message.success(response.msg)
          })
          this.pageList[index].Post.IsPublished = 0
        } else {
          publishPost(param).then(response => {
            this.$message.success(response.msg)
          })
          this.pageList[index].Post.IsPublished = 1
        }
        // this.getCount()
        // this.currentChangePage(this.currentPage)
      }
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
      fetchPostList(param).then(response => {
        var datas = response.data
        this.pageList = datas
        for (var i = 0; i < this.pageList.length; i++) {
          var date = new Date(Date.parse(this.pageList[i].Post.CreatedAt))
          this.pageList[i].Post.CreatedAt = date.toLocaleString('chinese', { hour12: false })
          this.pageList[i].Post.CategoryName = this.getCategoryNameById(this.pageList[i].Post.CategoryId)
          if (this.pageList[i].Post.IsDelete === 1) {
            this.pageList[i].Post.State = '已删除'
          } else if (this.pageList[i].Post.IsDelete === 0 && this.pageList[i].Post.IsPublished === 0) {
            this.pageList[i].Post.State = '草稿'
          } else if (this.pageList[i].Post.IsDelete === 0 && this.pageList[i].Post.IsPublished === 1) {
            this.pageList[i].Post.State = '已发布'
          }
        }
      })
    },
    getCategoryNameById(category_id) {
      for (var i = 0; i < this.categories.length; i++) {
        if (this.categories[i].Id === category_id) {
          return this.categories[i].Name
        }
      }
      return ''
    },
    // 重置表单信息
    resetForm(formName) {
      if (this.$refs[formName] !== undefined) {
        this.$refs[formName].resetFields()
      }
    },
    // 显示添加标签的抽屉
    handleOpenAddTagDraw() {
      this.drawOptions.addTagVisible = true
      this.resetForm('addCategoryForm')
    },
    // 添加标签事件
    handleTagAdd() {
      this.$refs.addTagForm.validate((valid) => {
        if (valid) {
          this.drawOptions.addTagBtnLoading = true
          var data = { 'name': this.addTagForm.name, 'description': this.addTagForm.description }
          addTag(data).then(response => {
            if (response.code === 100) {
              this.$message.success(response.msg)
              // this.getCount()
              // this.currentChangePage(this.currentPage)
              this.fetchTagData()
              // this.addForm = {}
              this.drawOptions.addTagVisible = false
            }
          })
          this.drawOptions.addTagBtnLoading = false
        }
      })
    },
    // 显示添加分类抽屉
    handleOpenAddCategoryDraw() {
      this.drawOptions.addVisible = true
      this.resetForm('addCategoryForm')
    },
    // 添加分类事件
    handleCategoryAdd() {
      this.$refs.addCategoryForm.validate((valid) => {
        if (valid) {
          this.drawOptions.addBtnLoading = true
          var data = { 'name': this.addCategoryForm.name, 'description': this.addCategoryForm.description }
          addCategory(data).then(response => {
            if (response.code === 100) {
              this.$message.success(response.msg)
              // this.getCount()
              // this.currentChangePage(this.currentPage)
              this.fetchCategoryData()
              // this.addForm = {}
              this.drawOptions.addVisible = false
            }
          })
          this.drawOptions.addBtnLoading = false
        }
      })
    },
    // 标签改变事件
    selectTrigger(formName) {
      this[`${formName}`].selectTagIds.forEach((id, i) => {
        if (typeof (id) === 'string') {
          addTag({ name: id })
            .then(res => {
              // 用 $set 动态修改标签列表，避免视图不更新
              this.$set(this.tags, this.tags.length, res.data)
              this[`${formName}`].selectTagIds[i] = res.data.ID
            })
            .catch(() => {
            })
        }
      })
    },
    getHtml() {
      this.html = this.$refs.markdownEditor.getHtml()
      this.dialogOptions.displayHTML = true
    },
    // 添加文章事件
    handleRowEdit() {
      // 获取编辑器组件中的文本内容
      this.editForm.content = this.$refs.markdownEditor.getHtml()
      // 获取 user_id
      this.editForm.author_id = Number(this.$store.state.user.userId)
      // 将数组转换成字符串
      this.editForm.tag_ids = this.editForm.selectTagIds.join()
      // 校验表单
      this.$refs.editForm.validate((valid) => {
        if (valid) {
          this.dialogOptions.addBtnLoading = true
          // setTimeout(() => {
          updatePost(this.editForm)
            .then(res => {
              this.$message.success(res.msg)
              this.dialogOptions.addVisible = false
              // this.fetchPageData()
            })
            .catch(() => {
            })
          this.dialogOptions.addBtnLoading = false
          this.dialogOptions.editVisible = false
          this.currentChangePage(1)
          this.getCount()
          this.fetchCategoryData()
          // }, 300)
        }
      })
    }

  }
}
</script>

<style scoped>

</style>
