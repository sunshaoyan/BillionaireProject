import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Dispatch, IRootState } from '@/store'
import { Button, Icon, Card, Radio, Affix } from 'antd'
import { COLOR_POOL } from '@/config/constants'
import forward from '@/assets/images/forward.png'
import backward from '@/assets/images/backward.png'
import left from '@/assets/images/left.png'
import right from '@/assets/images/right.png'
import turnLeft from '@/assets/images/turnLeft.png'
import turnRight from '@/assets/images/turnRight.png'
import leftHandStretch from '@/assets/images/leftHandStretch.png'
import rightHandStretch from '@/assets/images/rightHandStretch.png'
import leftHandFold from '@/assets/images/leftHandFold.png'
import rightHandFold from '@/assets/images/rightHandFold.png'
import leftHandUp from '@/assets/images/leftHandUp.png'
import rightHandUp from '@/assets/images/rightHandUp.png'
import playerOut from '@/assets/images/playerOut.png'
import playerIn from '@/assets/images/playerIn.png'

import './index.less'

interface IWorkspaceProps
  extends Partial<ReturnType<typeof mapState>>,
    Partial<ReturnType<typeof mapDispatch>> {}

const mapState = (state: IRootState) => ({
  scopeList: state.workspace.scopeList
})

const mapDispatch = (dispatch: Dispatch) => ({
  sendToCar: dispatch.workspace.sendToCar,
  updateList: dispatch.workspace.updateList,
  updateState: dispatch.workspace.updateState
})

class Workspace extends Component<IWorkspaceProps> {
  container: any
  constructor(props) {
    super(props)
    this.container = null
  }

  handleUpdateScope = () => {
    this.props.updateList(null)
  }

  handleDelete = index => {
    const { scopeList } = this.props
    scopeList.splice(index, 1)
    this.props.updateState({
      scopeList: [...scopeList]
    })
  }

  handleAddCondition = index => {
    const conditions = [
      'left_hand_stretch',
      'right_hand_stretch',
      'left_hand_fold',
      'right_hand_fold',
      'left_hand_up',
      'right_hand_up',
      'player_in',
      'player_out'
    ]
    const { scopeList } = this.props
    const scope = scopeList[index]
    scope.logic = 'or'
    conditions.splice(conditions.indexOf(scope.condition1), 1)
    scope.condition2 = conditions[0]
    this.props.updateState({
      scopeList: [...scopeList]
    })
  }

  handleChange = (e, index, part) => {
    const { value } = e.target
    console.log(value, index)
    const { scopeList } = this.props
    const scope = scopeList[index]
    switch (part) {
      case 'type':
        scope.type = value
        break
      case 'condition1':
        scope.condition1 = value
        break
      case 'logic':
        scope.logic = value
        break
      case 'condition2':
        scope.condition2 = value
        break
      case 'action':
        scope.action = value
        break
    }
    this.props.updateState({
      scopeList: [...scopeList]
    })
  }

  handleSendToCar = () => {
    this.props.sendToCar(null)
  }

  render() {
    const { scopeList } = this.props
    return (
      <div
        className="editor"
        ref={node => {
          this.container = node
        }}
      >
        <div className="submit">
          <Affix target={() => this.container}>
            <Button type="primary" onClick={this.handleSendToCar}>
              Send to car
            </Button>
          </Affix>
        </div>
        <div className="list">
          {scopeList.map((scope, index) => {
            const { type, condition1, condition2, logic, action } = scope
            let colorIndex = index * 3
            if (colorIndex + 3 > COLOR_POOL.length - 1) {
              colorIndex = 0
            }
            const cardBorder = COLOR_POOL[colorIndex]
            const conditionBorder = COLOR_POOL[colorIndex + 1]
            const actionBorder = COLOR_POOL[colorIndex + 2]
            return (
              <Card
                key={index}
                bordered={false}
                style={{
                  border: `3px solid ${cardBorder}`
                }}
                extra={
                  <Button
                    size="small"
                    type="danger"
                    shape="circle"
                    icon="close"
                    onClick={() => this.handleDelete(index)}
                  />
                }
              >
                <div className="condition-wrapper">
                  <div className="type">
                    <Radio.Group
                      buttonStyle="solid"
                      onChange={e => this.handleChange(e, index, 'type')}
                      value={type}
                    >
                      <Radio.Button value={'if'}>if</Radio.Button>
                      <Radio.Button value={'while'}>while</Radio.Button>
                    </Radio.Group>
                  </div>
                  <div
                    className="condition"
                    style={{
                      border: `2px dashed ${conditionBorder}`
                    }}
                  >
                    <Radio.Group
                      buttonStyle="solid"
                      onChange={e => this.handleChange(e, index, 'condition1')}
                      value={condition1}
                    >
                      <Radio value={'left_hand_stretch'}>
                        <img src={leftHandStretch} />
                      </Radio>
                      <Radio value={'right_hand_stretch'}>
                        <img src={rightHandStretch} />
                      </Radio>
                      <Radio value={'left_hand_fold'}>
                        <img src={leftHandFold} />
                      </Radio>
                      <Radio value={'right_hand_fold'}>
                        <img src={rightHandFold} />
                      </Radio>
                      <Radio value={'left_hand_up'}>
                        <img src={leftHandUp} />
                      </Radio>
                      <Radio value={'right_hand_up'}>
                        <img src={rightHandUp} />
                      </Radio>
                      <Radio value={'player_in'}>
                        <img className="player-in" src={playerIn} />
                      </Radio>
                      <Radio value={'player_out'}>
                        <img className="player-out" src={playerOut} />
                      </Radio>
                    </Radio.Group>
                    {logic !== 'none' ? (
                      <>
                        <div className="logic">
                          <Radio.Group
                            buttonStyle="solid"
                            onChange={e => this.handleChange(e, index, 'logic')}
                            value={logic}
                            style={{ textAlign: 'center' }}
                          >
                            <Radio.Button value={'or'}>or</Radio.Button>
                            <Radio.Button value={'and'}>and</Radio.Button>
                          </Radio.Group>
                        </div>
                        <Radio.Group
                          buttonStyle="solid"
                          onChange={e =>
                            this.handleChange(e, index, 'condition2')
                          }
                          value={condition2}
                        >
                          <Radio
                            value={'left_hand_stretch'}
                            disabled={condition1 === 'left_hand_stretch'}
                          >
                            <img src={leftHandStretch} />
                          </Radio>
                          <Radio
                            value={'right_hand_stretch'}
                            disabled={condition1 === 'right_hand_stretch'}
                          >
                            <img src={rightHandStretch} />
                          </Radio>
                          <Radio
                            value={'left_hand_fold'}
                            disabled={condition1 === 'left_hand_fold'}
                          >
                            <img src={leftHandFold} />
                          </Radio>
                          <Radio
                            value={'right_hand_fold'}
                            disabled={condition1 === 'right_hand_fold'}
                          >
                            <img src={rightHandFold} />
                          </Radio>
                          <Radio
                            value={'left_hand_up'}
                            disabled={condition1 === 'left_hand_up'}
                          >
                            <img src={leftHandUp} />
                          </Radio>
                          <Radio
                            value={'right_hand_up'}
                            disabled={condition1 === 'right_hand_up'}
                          >
                            <img src={rightHandUp} />
                          </Radio>
                          <Radio
                            value={'player_in'}
                            disabled={condition1 === 'player_in'}
                          >
                            <img className="player-in" src={playerIn} />
                          </Radio>
                          <Radio
                            value={'player_out'}
                            disabled={condition1 === 'player_out'}
                          >
                            <img className="player-out" src={playerOut} />
                          </Radio>
                        </Radio.Group>
                      </>
                    ) : (
                      <div style={{ textAlign: 'center' }}>
                        <Button
                          type="dashed"
                          block={true}
                          onClick={() => this.handleAddCondition(index)}
                          style={{ width: '60%' }}
                        >
                          <Icon type="plus" />
                          Add Condition
                        </Button>
                      </div>
                    )}
                  </div>
                </div>
                <div
                  className="action"
                  style={{
                    border: `2px dashed ${actionBorder}`
                  }}
                >
                  <Radio.Group
                    buttonStyle="solid"
                    onChange={e => this.handleChange(e, index, 'action')}
                    value={action}
                  >
                    <Radio value={'move_left'}>
                      <img src={left} alt="move left" />
                    </Radio>
                    <Radio value={'move_right'}>
                      <img src={right} alt="move right" />
                    </Radio>
                    <Radio value={'move_forward'}>
                      <img src={forward} alt="move forward" />
                    </Radio>
                    <Radio value={'move_backward'}>
                      <img src={backward} alt="move backward" />
                    </Radio>
                    <Radio value={'turn_left'}>
                      <img src={turnLeft} alt="turn left" />
                    </Radio>
                    <Radio value={'turn_right'}>
                      <img src={turnRight} alt="turn right" />
                    </Radio>
                  </Radio.Group>
                </div>
              </Card>
            )
          })}
        </div>

        <Button type="dashed" block={true} onClick={this.handleUpdateScope}>
          <Icon type="plus" />
          Add Program
        </Button>
      </div>
    )
  }
}

// export default Form.create<IFrameControlProps>()(
//   connect(
//     mapState,
//     mapDispatch
//   )(FrameControlContainer)
// )

export default connect(
  mapState,
  mapDispatch
)(Workspace)
