import React from "react";
import "../style/CourseCard.scss";

type CourseCardProps = {
  title: string;
  subtitle: string;
  description: string;
  courseUrl: string;
};

export class CourseCard extends React.Component<CourseCardProps> {
  render() {
    return (
      <div
        className="courseCard"
        onClick={(e) => (document.location.href = this.props.courseUrl)}
      >
        <p className="cardTitle">{this.props.title}</p>
        <p className="cardSubtitle">{this.props.subtitle}</p>
        <p className="cardDescription">{this.props.description}</p>
      </div>
    );
  }
}
