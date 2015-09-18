kd = require 'kd'
React = require 'kd-react'
Link  = require 'app/components/common/link'


module.exports = class PublicChannelLink extends React.Component

  @defaultProps = { to: null }


  ###
   * Allows @props.to to be either a channel thread or a channel itself.
   *
   * @param {string} key - key to be read from channel instance.
   * @return {*} value
  ###
  channel: (key) ->
    return  unless @props.to

    if @props.to.has 'channel'
    then @props.to?.getIn ['channel', key]
    else @props.to.get key


  render: ->
    <Link {...@props}
      className={kd.utils.curry "PublicChannelLink", @props.className}
      href="/Channels/#{@channel('name') ? '#'}">
      {@props.children}
    </Link>
