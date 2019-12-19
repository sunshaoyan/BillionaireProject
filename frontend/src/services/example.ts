const exampleJSON = {
  config: [
    {
      type: 'if', // or 'while'
      condition1: 'left_hand_stretch', // can be one of: '{left/right}_hand_{stretch/fold/up}' or 'player_{in/out}'
      condition2: 'right_hand_fold', //  same as condition1, but can also be empty string ''
      logic: 'or', // can be 'and', 'or', 'none'(representing that only condition1 is in consideration)
      action: 'move_left' // can be one of 'move_{ left / right /forward /backward }' or 'turn_{left / right }'
    }
  ]
}
