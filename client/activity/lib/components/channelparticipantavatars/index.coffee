kd                   = require 'kd'
React                = require 'kd-react'
Avatar               = require 'app/components/profile/avatar'
immutable            = require 'immutable'
classnames           = require 'classnames'
ProfileLinkContainer = require 'app/components/profile/profilelinkcontainer'


module.exports = class ChannelParticipantAvatars extends React.Component

  constructor: (props) ->

    super

    @state = { addNewParticipantMode: no }

  PREVIEW_COUNT = 0
  MAX_PREVIEW_COUNT = 4

  @defaultProps =
    channelThread : null
    participants  : null


  getPreviewCount: ->

    { participants } = @props

    diff = participants.size is MAX_PREVIEW_COUNT

    PREVIEW_COUNT = switch
      when diff = 0 then MAX_PREVIEW_COUNT
      when diff < 0 then participants.size
      else MAX_PREVIEW_COUNT - 1


  onNewParticipantButtonClick: ->

    if @state.addNewParticipantMode is yes
    then @setState addNewParticipantMode: no
    else @setState addNewParticipantMode: yes


  onShowMoreParticipantButtonClick: ->

    console.log 'show more participants clicked'


  renderAvatars: ->

    return null  unless @props.participants

    { participants } = @props

    PREVIEW_COUNT = @getPreviewCount()

    participants.slice(0, PREVIEW_COUNT).toList().map (participant) ->

  renderNickname: (participant, isNicknameVisible)->

    return  if isNicknameVisible is no

    nickname = participant.getIn ['profile', 'nickname']
    <span>{nickname}</span>


      <div key={participant.get 'id'} className="ChannelParticipantAvatars-singleBox">
        <ProfileLinkContainer account={participant.toJS()}>
          <Avatar
            className="ChannelParticipantAvatars-avatar"
            width={30}
            account={participant.toJS()}
            height={30} />
        </ProfileLinkContainer>
      </div>


  renderMoreCount: ->

    return null  unless @props.participants

    moreCount = @props.participants.size - PREVIEW_COUNT

    return null  unless moreCount > 0

    moreCount = Math.min moreCount, 99

    <div className="ChannelParticipantAvatars-singleBox">
      <div className="ChannelParticipantAvatars-moreCount" onClick={@bound "onShowMoreParticipantButtonClick"}>
        {moreCount}+
      </div>
    </div>


  getAddNewParticipantButtonClassNames: -> classnames
    'ChannelParticipantAvatars-newParticipantBox': yes
    'cross': @state.addNewParticipantMode


  renderNewParticipantButton: ->

    <div className="ChannelParticipantAvatars-singleBox" onClick={@bound "onNewParticipantButtonClick"}>
      <div className={@getAddNewParticipantButtonClassNames()}></div>
    </div>


  getNewParticipantInputClassNames: -> classnames
    'ChannelParticipantInput': yes
    'slide-down': @state.addNewParticipantMode


  renderAddNewParticipantInput: ->

    <div className={@getNewParticipantInputClassNames()}>
      <input placeholder="type a @username and hit enter" />
    </div>


  render: ->
    <div className="ChannelParticipantAvatars">
      {@renderAvatars()}
      {@renderMoreCount()}
      {@renderNewParticipantButton()}
      {@renderAddNewParticipantInput()}
    </div>

