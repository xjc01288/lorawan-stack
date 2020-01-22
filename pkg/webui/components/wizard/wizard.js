// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'

import PropTypes from '../../lib/prop-types'
import renderCallback from '../../lib/render-callback'

import WizardContext from './context'

const NEXT_STEP = 'NEXT_STEP'
const PREV_STEP = 'PREV_STEP'

const reducer = (state, action) => {
  switch (action.type) {
    case NEXT_STEP:
      return {
        ...state,
        step: state.step + 1,
      }
    case PREV_STEP:
      return {
        ...state,
        step: state.step - 1,
      }
    default:
      return state
  }
}

const Wizard = props => {
  const { initialStepNumber } = props

  const [state, dispatch] = React.useReducer(reducer, {
    step: initialStepNumber,
  })

  const next = React.useCallback(() => {
    dispatch({ type: NEXT_STEP })
  }, [])
  const prev = React.useCallback(() => {
    dispatch({ type: PREV_STEP })
  }, [])

  const context = React.useMemo(
    () => ({
      ...state,
      next,
      prev,
    }),
    [next, prev, state],
  )

  return (
    <WizardContext.Provider value={context}>
      {renderCallback(props, context)}
    </WizardContext.Provider>
  )
}

Wizard.propTypes = {
  initialStepNumber: PropTypes.number,
}

Wizard.defaultProps = {
  initialStepNumber: 1,
}

export default Wizard
