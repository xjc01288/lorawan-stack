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
import { storiesOf } from '@storybook/react'

import Wizard from '.'

const { Stepper } = Wizard

// eslint-disable-next-line react/prop-types
const Steps = ({ next, prev, step }) => {
  return (
    <div>
      <Stepper status="current">
        <Stepper.Step title="Step 1" stepNumber={1} />
        <Stepper.Step title="Step 2" stepNumber={2} />
        <Stepper.Step title="Step 3" stepNumber={3} />
        <Stepper.Step title="Step 4" stepNumber={4} />
      </Stepper>
      <Wizard.Step stepNumber={1}>
        <span>Step 1 content</span>
      </Wizard.Step>
      <Wizard.Step stepNumber={2}>
        <span>Step 2 content</span>
      </Wizard.Step>
      <Wizard.Step stepNumber={3}>
        <span>Step 3 content</span>
      </Wizard.Step>
      <Wizard.Step stepNumber={4}>
        <span>Step 4 content</span>
      </Wizard.Step>
      <div style={{ display: 'flex', justifyContent: 'space-between' }}>
        <button disabled={step <= 1} onClick={prev}>
          Previous
        </button>
        <button disabled={step >= 4} onClick={next}>
          Next
        </button>
      </div>
    </div>
  )
}

const stepsRenderer = wizardProps => <Steps {...wizardProps} />

storiesOf('Wizard', module).add('Default', () => <Wizard render={stepsRenderer} />)
