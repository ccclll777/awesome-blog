<script>
export default {
  created() {
    const { params, query } = this.$route
    const { path } = params
    console.log('123131312')
    this.$router.replace({ path: '/' + path, query })
  },
  render: function(h) {
    return h() // avoid warning message
  }
}
</script>
