new Vue({
  el: '#app',

  data: {

    memos: [],

    memoDate: new Date(),

    memoContent: '',

    memoStatus: 'incomplete',

    assignedMember: '',
    optionMember: [
      { id: 1, name: 'Mike' },
      { id: 2, name: 'Tom' },
      { id: 3, name: 'Ann' }
    ],

    isEntered: false,

    customFormat: 'yyyy-MM-dd'
  },

  components: {
    'vuejs-datepicker': vuejsDatepicker
  },

  computed: {
    computedMemos () {
      return this.memos
    },

    validate () {
      var isEnteredMemoContent = 0 < this.memoContent.length
      this.isEntered = isEnteredMemoContent
      return isEnteredMemoContent
    },
  },

  created: function () {
    this.fetchAllMemos()
  },



  methods: {
    fetchAllMemos() {
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
    },

    addMemo() {
      const params = new URLSearchParams();
      this.memoDate = moment(this.memoDate).format('YYYY-MM-DD')
      params.append('memoDate', this.memoDate)
      params.append('memoContent', this.memoContent)
      params.append('memoStatus', this.memoStatus)

      axios.post('/addMemo', params)
      .then(response => {
        if(response.status != 200) {
          throw new Error('Response Error')
        }
        else {
          this.fetchAllMemos()

          this.initInputValue()
        }
      })
    },

    deleteMemo(memo) {
      const params = new URLSearchParams();
      params.append('memoID', memo.id)

      axios.post('/deleteMemo', params)
      .then(response => {
        if(response.status != 200) {
          throw new Error('Response Error')      
        }
        else {
          this.fetchAllMemos()
        }
      })
    },

    initInputValue () {
      this.memoContent = ''
      this.memoStatus = ''
    },

  }
})