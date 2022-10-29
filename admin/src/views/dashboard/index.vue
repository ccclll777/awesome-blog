<template>
  <div class="dashboard-editor-container">
    <el-row>
      <el-col :span="8">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文章</span>
            <el-button
              size="medium"
              icon="el-icon-notebook-2"
              style="float: right; padding: 3px 0"
              type="text"
              @click="$router.push('/post')"
            />
          </div>
          <div class="text item"><h2 style="color: #2f74ff" class="card-num">{{ indexData.article_count }}</h2></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>评论</span>
            <el-button
              size="medium"
              icon="el-icon-s-comment"
              style="float: right; padding: 3px 0"
              type="text"
              @click="$router.push('/comment')"
            />
          </div>
          <div class="text item"><h2 style="color: #4dc820" class="card-num">{{ indexData.comment_count }}</h2></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>访问量</span>
            <el-button size="medium" icon="el-icon-view" style="float: right; padding: 3px 0" type="text" />
          </div>
          <div class="text item"><h2 style="color: slategrey" class="card-num">0</h2></div>
        </el-card>
      </el-col>
      <el-col :span="24">
        <el-tabs class="box-card" type="border-card">
          <el-tab-pane label="最近文章">
            <el-timeline>
              <p v-if="indexData.latest_articles.length === 0" class="no-tip">暂无文章</p>
              <el-timeline-item
                v-for="item in indexData.latest_articles"
                :key="item.ID"
                :timestamp="formatDate(item.CreatedAt)"
                placement="top"
              >
                {{ item.title }}
              </el-timeline-item>
            </el-timeline>
          </el-tab-pane>
          <el-tab-pane label="最近评论">
            <el-timeline>
              <p v-if="indexData.latest_comments.length === 0" class="no-tip">暂无评论</p>
              <el-timeline-item
                v-for="item in indexData.latest_comments"
                :key="item.ID"
                :timestamp="formatDate(item.CreatedAt)"
                placement="top"
              >
                {{ item.content }}
              </el-timeline-item>
            </el-timeline>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
  </div>
</template>

<script>

export default {
  name: 'DashboardAdmin',
  components: {
  },
  data() {
    return {
      indexData: {
        article_count: 0,
        comment_count: 0,
        latest_articles: [],
        latest_comments: []
      },
      fmt: 'yyyy-MM-dd',
      currentDate: new Date()
    }
  },
  methods: {

  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
