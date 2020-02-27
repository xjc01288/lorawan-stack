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
import classnames from 'classnames'

import Event from '../..'
import Message from '../../../../lib/components/message'
import Icon from '../../../icon'
import PropTypes from '../../../../lib/prop-types'
import { getEntityId } from '../../../../lib/selectors/id'
import { warn } from '../../../../lib/log'
import style from './crud.styl'
import { formatMessageData, getEventActionByName, getErrorEvent } from '..'

class CRUDEvent extends React.PureComponent {
  static propTypes = {
    className: PropTypes.string,
    event: PropTypes.event.isRequired,
    expandedClassName: PropTypes.string,
    overviewClassName: PropTypes.string,
    widget: PropTypes.bool,
  }

  static defaultProps = {
    className: undefined,
    expandedClassName: undefined,
    overviewClassName: undefined,
    widget: false,
  }

  render() {
    const { className, event, widget, overviewClassName, expandedClassName } = this.props

    const entityId = getEntityId(event.identifiers[0])
    const eventAction = getEventActionByName(event.name)
    const data = formatMessageData(event.data)
    const isError = getErrorEvent(event.data)
    let icon = null

    if (isError) {
      icon = <Icon icon="error" className={style.error} />
    } else if (eventAction === 'create') {
      icon = <Icon icon="event_create" className={style.create} />
    } else if (eventAction === 'delete') {
      icon = <Icon icon="event_delete" className={style.delete} />
    } else if (eventAction === 'update') {
      icon = <Icon icon="event_update" />
    } else {
      warn(`Unknown event name: ${event.name}`)
      icon = <Icon icon="event" />
    }

    const content = (
      <Message
        className={classnames({ [style.error]: isError })}
        content={{ id: `event:${event.name}` }}
      />
    )

    return (
      <Event
        className={className}
        overviewClassName={overviewClassName}
        expandedClassName={expandedClassName}
        icon={icon}
        time={event.time}
        emitter={entityId}
        content={content}
        widget={widget}
        data={data}
      />
    )
  }
}

export default CRUDEvent
