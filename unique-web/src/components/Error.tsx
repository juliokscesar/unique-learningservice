import React from "react";
import { RouteComponentProps } from "react-router-dom";

interface ErrorProps extends RouteComponentProps<any> {}

type ErrorState = {
  err: string;
};

export class Error extends React.Component<ErrorProps, ErrorState> {
  state: ErrorState = {
    err: "",
  };

  componentDidMount() {
    const params = new URLSearchParams(this.props.location.search);
    const errParam = params.get("err");

    this.setState({
      err: errParam === null ? "Error page" : errParam
    });
  }

  render() {
    return (
      <div className="errorPage">
        <h1>Sorry, something wrong occurred.</h1>

        <p>{this.state.err}</p>
      </div>
    );
  }
}
