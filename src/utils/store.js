import { createStore } from 'vuex'

// 创建一个新的 store 实例
const store = createStore({
    state () {
        return {
            count: 0,
            name: ''  // 添加 name 到 state
        }
    },
    mutations: {
        increment (state) {
            state.count++
        },
        setName (state, name) {  // 添加一个 mutation 来设置 name
            state.name = name
        }
    },
    getters: {
        getName: state => state.name  // 添加一个 getter 来获取 name
    }
})

export {store}
