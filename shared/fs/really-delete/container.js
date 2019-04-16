// @flow
import * as FsGen from '../../actions/fs-gen'
import * as RouteTreeGen from '../../actions/route-tree-gen'
import * as Container from '../../util/container'
import ReallyDelete from '.'
import {anyWaiting} from '../../constants/waiting'
import * as Types from '../../constants/types/fs'
import * as Constants from '../../constants/fs'

type OwnProps = Container.RouteProps<{path: Types.Path}, {}>

const mapStateToProps = state => {
  return {
    _deleting: (path: Types.Path) => anyWaiting(state, Constants.deleteFolderWaitingKey(path)),
    title: 'Confirmation',
  }
}

const mapDispatchToProps = dispatch => ({
  _onDelete: (path: Types.Path) => {
    dispatch(FsGen.createDeleteFile({path}))
    // TODO: I think I should be closing this using the waiting key instead,
    //  but I'm not sure how that works.
    dispatch(RouteTreeGen.createNavigateUp())
  },
  onBack: () => dispatch(RouteTreeGen.createNavigateUp()),
})

const mergeProps = (stateProps, dispatchProps, ownProps) => {
  const path = Container.getRouteProps(ownProps, 'path')
  return {
    _deleting: stateProps._deleting(path),
    onBack: stateProps._deleting(path) ? () => {} : dispatchProps.onBack,
    onDelete: () => dispatchProps._onDelete(path),
    path: path,
    title: stateProps.title,
  }
}

export default Container.compose(
  Container.connect<OwnProps, _, _, _, _>(
    mapStateToProps,
    mapDispatchToProps,
    mergeProps
  ),
  Container.safeSubmit(['onDelete'], ['_deleting'])
)(ReallyDelete)
