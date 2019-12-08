import React, { Component } from 'react'

export default class Button extends Component {
    
    constructor(props) {
        super(props)
        this.props.eventHandler = this.eventHandler.bind(this)
        this.state = {
             buttonName : this.props.buttonName
        }
    }
    
    render() {
        return (
            <div>
                <button onClick={this.props.eventHandler}>{this.state.buttonName}</button>
            </div>
        )
    }
}
