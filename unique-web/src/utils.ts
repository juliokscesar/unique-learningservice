import { UNIQUE_BASE_TITLE } from "./constants";

export const toggleElementById = (id: string, toggle: boolean) => {
  const element = document.getElementById(id) as HTMLInputElement;
  element.disabled = !toggle;
};

const emailRegex =
  /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
export const validateEmail = (email: string) => {
  return emailRegex.test(email);
};

export const capitalizeFirstLetter = (str: string) => {
  return str.charAt(0).toLocaleUpperCase() + str.slice(1);
};

export const setTitle = (title: string) => {
  document.title = UNIQUE_BASE_TITLE + title;
};
