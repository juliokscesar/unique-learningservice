import React from "react";
import { RouteComponentProps } from "react-router-dom";
import { UNIQUE_BASE_TITLE } from "../constants";

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
      err: errParam === null ? "Error page" : errParam,
    });

    document.title = UNIQUE_BASE_TITLE + "Error";
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
