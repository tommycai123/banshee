import React from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router'

import Paper from 'material-ui/lib/paper'
import Toolbar from 'material-ui/lib/toolbar/toolbar'
import ToolbarGroup from 'material-ui/lib/toolbar/toolbar-group'
import ToolbarSeparator from 'material-ui/lib/toolbar/toolbar-separator'
import ToolbarTitle from 'material-ui/lib/toolbar/toolbar-title'
import RaisedButton from 'material-ui/lib/raised-button'
import IconButton from 'material-ui/lib/icon-button'
import List from 'material-ui/lib/lists/list'
import ListItem from 'material-ui/lib/lists/list-item'
import FlatButton from 'material-ui/lib/flat-button'
import TextField from 'material-ui/lib/text-field'
import Dialog from 'material-ui/lib/dialog'
import Checkbox from 'material-ui/lib/checkbox'

import ActionDelete from 'material-ui/lib/svg-icons/action/delete'

import { actions as projectDetailActions } from '../redux/modules/projectDetail'

import ruleStyles from '../styles/rule.scss'

const mapStateToProps = (state) => ({
  project: state.projectDetail.project,
  rules: state.projectDetail.rules,
  id: state.router.path.split('/')[1],
  ruleOpen: state.projectDetail.ruleOpen,
  submitDisabled: state.projectDetail.submitDisabled,
  patternErrorText: state.projectDetail.patternErrorText,
  onTrendingUp: state.projectDetail.onTrendingUp,
  onTrendingDown: state.projectDetail.onTrendingDown,
  onValueGt: state.projectDetail.onValueGt,
  onValueLt: state.projectDetail.onValueLt,
  onTrendingUpAndValueGt: state.projectDetail.onTrendingUpAndValueGt,
  onTrendingDownAndValueLt: state.projectDetail.onTrendingDownAndValueLt,
  thresholdMax: state.projectDetail.thresholdMax,
  thresholdMin: state.projectDetail.thresholdMin,
  trustline: state.projectDetail.trustline,
  thresholdMaxErrorText: state.projectDetail.thresholdMaxErrorText,
  thresholdMinErrorText: state.projectDetail.thresholdMinErrorText,
  trustlineErrorText: state.projectDetail.trustlineErrorText
})

export class RuleList extends React.Component {
  static propTypes = {
    getProjectById: React.PropTypes.func.isRequired,
    getAllRules: React.PropTypes.func.isRequired,
    ruleDialogClose: React.PropTypes.func.isRequired,
    id: React.PropTypes.string.isRequired,
    ruleOpen: React.PropTypes.bool.isRequired,
    addRule: React.PropTypes.func.isRequired,
    project: React.PropTypes.object.isRequired,

    handlePatternChange: React.PropTypes.func.isRequired,
    handleCheck: React.PropTypes.func.isRequired,
    handleInput: React.PropTypes.func.isRequired,

    submitDisabled: React.PropTypes.bool.isRequired,
    onTrendingUp: React.PropTypes.bool.isRequired,
    onTrendingDown: React.PropTypes.bool.isRequired,
    onValueGt: React.PropTypes.bool.isRequired,
    onValueLt: React.PropTypes.bool.isRequired,
    onTrendingUpAndValueGt: React.PropTypes.bool.isRequired,
    onTrendingDownAndValueLt: React.PropTypes.bool.isRequired,
    thresholdMax: React.PropTypes.number,
    thresholdMin: React.PropTypes.number,
    trustline: React.PropTypes.number,

    patternErrorText: React.PropTypes.string.isRequired,
    thresholdMaxErrorText: React.PropTypes.string.isRequired,
    thresholdMinErrorText: React.PropTypes.string.isRequired,
    trustlineErrorText: React.PropTypes.string.isRequired
  }

  componentDidMount () {
    let id = this.props.id
    this.props.getProjectById(id)
    this.props.getAllRules(id)
  }

  render () {
    const styles = {
      leftBtn: {
        marginRight: 0
      },
      toolbar: {
        backgroundColor: '#fff'
      },
      separator: {
        margin: '0 16px 0 0'
      },
      link: {
        textDecoration: 'none',
        color: '#444'
      },
      hover: {
        textDecoration: 'underline'
      }
    }

    const rightIconButton = (
      <IconButton tooltip='delete' tooltipPosition='top-right' >
        <ActionDelete />
      </IconButton>
    )

    const ruleActions = [
      <FlatButton
        label='Cancel'
        secondary
        onTouchTap={this.props.ruleDialogClose} />,
      <FlatButton
        label='Submit'
        primary
        form='form'
        disabled={this.props.submitDisabled}
        onTouchTap={this.props.addRule} />
    ]

    return (
      <Paper zDepth={1}>
        <Toolbar style={styles.toolbar}>
          <ToolbarGroup float='left'>
            <ToolbarTitle text={this.props.project.name} />
            <ToolbarSeparator style={styles.separator}/>
            <ToolbarTitle text='Rules' />
          </ToolbarGroup>
          <ToolbarGroup float='right'>
            <RaisedButton label='Edit Name' primary style={styles.leftBtn}/>
            <ToolbarSeparator />
            <RaisedButton label='Add Rule' primary/>
          </ToolbarGroup>
        </Toolbar>
        <List>
          <ListItem
            primaryText={<Link className={ruleStyles.link} to='/'> project name</Link>}
            secondaryText='Change your Google+ profile photo'
            rightIconButton={rightIconButton}/>
        </List>
        <Dialog
          title='Add Rule'
          actions={ruleActions}
          modal={false}
          open
          onRequestClose={this.props.ruleDialogClose}>
            <form id='form' onSubmit={this.props.addRule}>
              <div className={ruleStyles.row}>
                <label className={ruleStyles.label}>Pattern:</label>
                <div className={ruleStyles.rightPart}>
                  <TextField
                    className={ruleStyles.verticalAlign}
                    hintText='timer.count_ps.*'
                    onChange={this.props.handlePatternChange}
                    errorText={this.props.patternErrorText}/>
                </div>
              </div>

              <div className={ruleStyles.row}>
                <label className={ruleStyles.label}>Alerting:</label>
                <div className={ruleStyles.rightPart}>
                  <Checkbox
                    onCheck={(e, checked) => { this.props.handleCheck('onTrendingUp', checked) }}
                    label='On trending up'/>
                  <Checkbox
                    onCheck={(e, checked) => { this.props.handleCheck('onTrendingDown', checked) }}
                    label='On trending down'/>
                  <Checkbox
                    onCheck={(e, checked) => { this.props.handleCheck('onValueGt', checked) }}
                    label='On value >= thresholdMax'/>
                  <Checkbox
                    onCheck={(e, checked) => { this.props.handleCheck('onValueLt', checked) }}
                    label='On value <= thresholdMin'/>
                  <Checkbox
                    onCheck={(e, checked) => { this.props.handleCheck('onTrendingUpAndValueGt', checked) }}
                    label='On trending up and value >= thresholdMax'/>
                  <Checkbox
                    onCheck={(e, checked) => { this.props.handleCheck('onTrendingDownAndValueLt', checked) }}
                    label='On trending down and value <= thresholdMin'/>
                  <div>
                    <label>thresholdMax:</label>
                    <TextField
                      value={this.props.thresholdMax}
                      className={ruleStyles.smallField}
                      onChange={(e) => { this.props.handleInput('thresholdMax', e.target.value) }}
                      errorText={this.props.thresholdMaxErrorText}/>
                  </div>
                  <div>
                    <label>thresholdMin:</label>
                    <TextField
                      value={this.props.thresholdMin}
                      className={ruleStyles.smallField}
                      onChange={(e) => { this.props.handleInput('thresholdMin', e.target.value) }}
                      errorText={this.props.thresholdMinErrorText}/>
                  </div>
                </div>
              </div>

              <div className={ruleStyles.row}>
                <label className={ruleStyles.label}>Trustline:</label>
                <div className={ruleStyles.rightPart && ruleStyles.divVertical}>
                  <label>Don't alert me when value is less than</label>
                  <TextField
                      value={this.props.trustline}
                      className={ruleStyles.smallField}
                      onChange={(e) => { this.props.handleInput('trustline', e.target.value) }}
                      errorText={this.props.trustlineErrorText}/>
                </div>
              </div>

            </form>
        </Dialog>
      </Paper>
    )
  }
}

// export default RuleList
export default connect(mapStateToProps, projectDetailActions)(RuleList)
