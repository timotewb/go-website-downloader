import "./Download.css";
import { useRef, useState } from "react";
import * as App from "../../../wailsjs/go/main/App";

function Download() {
  const inputRef = useRef("");
  const [urlValue, setURLVaue] = useState("");
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    inputRef.current = e.target.value;
  };
  const handleBlur = (e: React.FocusEvent<HTMLInputElement, Element>) => {
    console.log("handleBlur");
    e.target.placeholder = "Enter URL";
  };
  const handleFocus = (e: React.FocusEvent<HTMLInputElement, Element>) => {
    console.log("handleFocus");
    e.target.placeholder = "";
  };
  const handleClick = () => {
    console.log(inputRef);
    App.FindURL(inputRef.current);
  };
  return (
    <>
      <div id="input">
        <input
          id="inputArea"
          placeholder="Enter URL"
          onFocus={(e) => {
            handleFocus(e);
          }}
          onBlur={(e) => {
            handleBlur(e);
          }}
          onChange={(e) => handleChange(e)}
        ></input>
        <div id="inputButton" onClick={() => handleClick()}>
          Find
        </div>
      </div>
    </>
  );
}

export default Download;
