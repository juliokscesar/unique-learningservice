import React from "react";
import * as $ from "jquery";

const getFromAPI = (apiUrl: string) => {
  $.ajax({
    url: apiUrl,
    type: "GET",
    dataType: "json",
    success: (res) => {
      console.log(res);
    },
    error: (err) => {
      console.log(err);
    }
  });
  return "Done!";
}

type TestApiProps = {
  urlTest: string;
};

export const TestApi = (props: TestApiProps) => {
  const result = getFromAPI(props.urlTest);

  console.log(result);

  return (
      <div>
          <h1>{ result }</h1>
      </div>
  )
}
