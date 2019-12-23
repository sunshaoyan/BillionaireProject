import { createModel } from '@rematch/core'
import * as services from '@/services/index'
import { message } from 'antd'

export const workspace = createModel({
  state: {
    scopeList: [
      {
        type: 'if', // or 'while'
        condition1: 'left_hand_stretch', // can be one of: '{left/right}_hand_{stretch/fold/up}' or 'player_{in/out}'
        condition2: '', //  same as condition1, but can also be empty string ''
        logic: 'none', // can be 'and', 'or', 'none'(representing that only condition1 is in consideration)
        action: 'move_left' // can be one of 'move_{ left / right /forward /backward }' or 'turn_{left / right }'
      }
    ]
  },

  reducers: {
    // 在这里写纯函数来改变 state
    updateState(state, payload) {
      return {
        ...state,
        ...payload
      }
    }
  },

  effects: () => ({
    async sendToCar(_, rootState) {
      const { scopeList } = rootState.workspace
      const data = await services.workspace.sendToCar(scopeList)
      if (data.code === 0) {
        message.success('send success')
      }
    },

    updateList(_payload, rootState) {
      const { scopeList } = rootState.workspace
      const defaultScope = {
        type: 'if',
        condition1: 'left_hand_stretch',
        condition2: '',
        logic: 'none',
        action: 'move_left'
      }
      scopeList.push(defaultScope)
      this.updateState({
        scopeList: [...scopeList]
      })
    }
  })
})
