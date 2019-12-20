import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Dispatch, IRootState } from '@/store'
import './index.less'

interface ICodePreviewProps
  extends Partial<ReturnType<typeof mapState>>,
    Partial<ReturnType<typeof mapDispatch>> {}

const mapState = (state: IRootState) => ({
  scopeList: state.workspace.scopeList
})

const mapDispatch = (dispatch: Dispatch) => ({
  sendToCar: dispatch.workspace.sendToCar
})

const logicMap = {
  or: '||',
  and: '&&'
}

class CodePreview extends Component<ICodePreviewProps> {
  constructor(props) {
    super(props)
  }

  render() {
    const { scopeList } = this.props
    return (
      <div className="code-preview">
        {scopeList.map((scope, index) => {
          const { type, condition1, condition2, logic, action } = scope
          const condition =
            logic !== 'none'
              ? `${condition1} ${logicMap[logic]} ${condition2}`
              : condition1
          return (
            <p key={index}>
              <code>
                {`${type} (${condition}) {
                  ${action}
                }`}
              </code>
            </p>
          )
        })}
      </div>
    )
  }
}

export default connect(
  mapState,
  mapDispatch
)(CodePreview)
