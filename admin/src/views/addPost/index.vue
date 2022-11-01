<template>
  <div class="errPage-container">
    <el-form ref="addForm" :model="addForm" :rules="addRules" label-width="80px" style="margin-top: 50px">
      <el-form-item label="标题" prop="title">
        <el-input v-model="addForm.title" size="small" type="text" />
      </el-form-item>
      <el-row>
        <el-col :span="8">
          <el-form-item label="分类" prop="category_id">
            <!--            <el-select size="small" v-model="addForm.category_id" clearable @clear="editForm.category_id=null"-->
            <el-select
              v-model="addForm.category_id"
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
              v-model="addForm.selectTagIds"
              size="small"
              filterable
              allow-create
              multiple
              :multiple-limit="3"
              clearable
              placeholder="请选择标签"
              @change="selectTrigger('addForm')"
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
            <el-input v-model="addForm.keywords" size="small" type="text" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="摘要" prop="description">
        <el-input v-model="addForm.description" :rows="5" type="textarea" />
      </el-form-item>
      <el-form-item label="内容" prop="content">
        <markdown-editor ref="markdownEditor" v-model="content" :options="{hideModeSwitch:true,previewStyle:'tab'}" height="200px" />
      </el-form-item>

      <el-form-item>
        <el-button
          type="success"
          :loading="dialogOptions.addDraftBtnLoading"
          @click="handleRowAdd(false)"
        >保存为草稿
        </el-button>
        <el-button
          type="primary"
          :loading="dialogOptions.addBtnLoading"
          @click="handleRowAdd(true)"
        >发布
        </el-button>
        <el-button @click="dialogOptions.addVisible=false">取消</el-button>
        <el-button style="margin-top:80px;" type="primary" icon="el-icon-document" @click="getHtml">
          内容预览
        </el-button>
      </el-form-item>
    </el-form>
    <!-- 添加文章弹窗 -->
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
  </div>
</template>

<script>
const content1 = ``
import MarkdownEditor from '@/components/MarkdownEditor'
import { addCategory, fetchAllCategory } from '@/api/category'
import { addTag, fetchAllTag } from '@/api/tag'
export default {
  name: 'Post',
  components: { MarkdownEditor },
  data() {
    return {
      html: '', // 将mark渲染成HTML，然后输出
      content: content1,
      addForm: {
        author_id: null,
        category_id: null,
        selectTagIds: [],
        tag_ids: '',
        is_published: true,
        title: '',
        description: '',
        content: '',
        md_content: '',
        keywords: ''
      },
      dialogOptions: {
        displayHTML: false,
        addBtnLoading: false,
        addDraftBtnLoading: false,
        addVisible: false,
        addFullScreen: true,
        editBtnLoading: false,
        editDraftBtnLoading: false,
        editVisible: false,
        editFullScreen: true,
        uploadVisible: false,
        uploadLoading: false
      },
      drawOptions: {
        addVisible: false,
        addBtnLoading: false,
        addTagVisible: false,
        addTagBtnLoading: false
      },
      addRules: {
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
      categories: [],
      tags: [],
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
  created() {
    this.fetchCategoryData()
    this.fetchTagData()
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
      // console.log(this.addForm.selectTagIds)
    },
    // 添加文章事件
    handleRowAdd(isPublished) {
      // 获取编辑器组件中的文本内容
      console.log('handleRowAdd')
      this.addForm.content = this.content
      console.log('this.content', this.content)
      this.addForm.md_content = this.$refs.markdownEditor.getHtml()
      console.log('this.html', this.addForm.md_content)
      // 获取 user_id
      this.addForm.author_id = Number(localStorage.getItem('uuid'))
      // 将数组转换成字符串
      this.addForm.tag_ids = this.addForm.selectTagIds.join()
      console.log('this.addForm.tag_ids', this.addForm.tag_ids)
      this.addForm.is_published = isPublished
      // 校验表单
      this.$refs.addForm.validate((valid) => {
        if (valid) {
          if (isPublished) {
            this.dialogOptions.addBtnLoading = true
          } else {
            this.dialogOptions.addDraftBtnLoading = true
          }
          // setTimeout(() => {
          //   addPost(this.addForm)
          //     .then(res => {
          //       this.$message.success(res.msg)
          //       this.dialogOptions.addVisible = false
          //       this.fetchPageData()
          //     })
          //     .catch(() => {
          //     })
          //   if (isPublished) {
          //     this.dialogOptions.addBtnLoading = false
          //   } else {
          //     this.dialogOptions.addDraftBtnLoading = false
          //   }
          // }, 300)
        }
      })
    },
    // 打开弹窗
    handleOpenDrawer(mode) {
      this.$refs.attachDrawer.openDrawer()
      this.mode = mode
    },
    getHtml() {
      this.html = this.$refs.markdownEditor.getHtml()
      this.dialogOptions.displayHTML = true
    }
  }
}
</script>

<style scoped>

</style>
