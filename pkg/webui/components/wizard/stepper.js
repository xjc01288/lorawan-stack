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

import Stepper from '../stepper'

import WizardContext from './context'

const WizardStepper = props => {
  const { step } = React.useContext(WizardContext)
  const { children, ...rest } = props

  return (
    <Stepper {...rest} currentStep={step}>
      {children}
    </Stepper>
  )
}

WizardStepper.propTypes = {
  children: PropTypes.oneOfType([PropTypes.node, PropTypes.arrayOf(PropTypes.node)]),
}

WizardStepper.defaultProps = {
  children: [],
}

WizardStepper.Step = Stepper.Step
WizardStepper.displayName = 'Wizard.Stepper'

export default WizardStepper
