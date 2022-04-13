import React from "react";
import axios from 'axios';

interface VersionInfoState {
  version: {version: string, build_timestamp: string, commit_hash: string} | null
}

export default class VersionInfo extends React.Component<{}, VersionInfoState> {
  componentDidMount() {
    axios.get('/api/v1/version')
      .then(res => {
        const version = res.data
        this.setState({version})
      })
  }

  render() {
    if(this.state == null) {
      return (
        <h3>...</h3>
      );
    } else {
      return (
        <h3><>flowflow {this.state.version?.version}</></h3>
      );
    }
  }
}