import "./Download.css";
import { useRef } from "react";

function Download() {
  const inputRef = useRef("");
  const handleChage = (e: React.ChangeEvent<HTMLInputElement>) => {
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
          onChange={(e) => handleChage(e)}
        ></input>
        <div id="inputButton">Find</div>
      </div>
    </>
  );
}

export default Download;
