import React from "react";

type ToastProps = {
  message: string;
  author: string;
};

export class Toast extends React.Component<ToastProps> {
  render() {
    return (
      <div>
        <p>{this.props.message}</p>
        <p style={{ fontStyle: "italic" }}>{this.props.author}</p>
      </div>
    );
  }
}
