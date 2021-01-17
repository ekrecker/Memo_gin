new Vue({
  el: '#app',

  data: {
    memos: [
    ],

    memoDate: '',

    memoContent: '',

    memoStatus: '',
  },
  computed: {
    computedMemos () {
      return this.memos
    },
  },

  created: function () {
    this.FetchAllMemos()
  },

  methods: {

    FetchAllMemos () {
      axios.get('/fetchAllMemos')
      .then(response => {
        if(response.status != 200) {
          throw new Error('Response Error')
        }
        else {
          var resultMemos = response.data

          this.memos = resultMemos
        }
      })
    }
  }
})