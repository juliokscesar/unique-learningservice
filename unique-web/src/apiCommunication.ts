import QueryString from "qs";
import { API_BASE_URI, API_AUTH_USER, API_AUTH_PASS } from "./constants";
import { encode } from "base-64";

const getRequestAPI = (urlReq: string) => {
  return fetch(API_BASE_URI + urlReq, {
    method: "GET",
    headers: {
      Authorization: "Basic " + encode(API_AUTH_USER + ":" + API_AUTH_PASS),
    },
  }).then((res) => res.json());
};

const postFormRequestAPI = async (urlReq: string, formData: any) => {
  const data = await fetch(API_BASE_URI + urlReq, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
      Authorization: "Basic " + encode(API_AUTH_USER + ":" + API_AUTH_PASS),
    },
    body: QueryString.stringify(formData),
  })
    .then((res) => res.json())
    .then((data) => JSON.stringify(data));

  return JSON.parse(data);
};

export const getUserFromId = (id: string) => {
  return getRequestAPI("user/" + id);
};

export const getUserFromPublicId = (pid: string) => {
  return getRequestAPI("user/profile/" + pid);
};

export const getCourseFromId = (cid: string) => {
  return getRequestAPI("course/" + cid);
};

export const getManyCoursesFromIds = (cids: string[]) => {
  return getRequestAPI("courses/" + String(cids));
};

export const registerUser = (formData: any) => {
  return postFormRequestAPI("user/register", formData);
};

export const loginUser = (formData: any) => {
  return postFormRequestAPI("user/login", formData);
};

export const changeUserPassword = (uid: string, formData: any) => {
  return postFormRequestAPI(
    "user/" + uid + "/settings/change/password",
    formData
  );
};

export const createCourse = (formData: any) => {
  return postFormRequestAPI("course/create", formData);
};
