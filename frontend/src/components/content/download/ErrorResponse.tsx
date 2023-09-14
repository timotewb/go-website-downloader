import { lib } from "../../../../wailsjs/go/models";
import "./shared.css";
import errorSVG from "../../../assets/images/error.svg";
import { useState } from "react";

function ErrorResponse(props: lib.ResponseType) {
  const [showErrorMessage, setShowErrorMessage] = useState(false);
  // identify which error message to use
  let message: string;
  switch (props.code) {
    case 0:
      message = "System error. Code 0 receieved.";
      break;
    case 1:
      message = "Please enter a valid URL (e.g. https://www.website.com/blog)";
      break;
    default:
      message = "System error. No error code receieved.";
      break;
  }

  const handleClick = () => {
    if (showErrorMessage) {
      setShowErrorMessage(false);
    } else {
      setShowErrorMessage(true);
    }
  };

  const ErrorMessageBox = () => {
    if (showErrorMessage) {
      return <div id="errorMessageBox">{props.message}</div>;
    } else {
      return <></>;
    }
  };

  return (
    <>
      <div id="messageArea">
        <div className="icon">
          <img src={errorSVG}></img>
        </div>
        <div className="message">{message}</div>
      </div>
      <div id="viewErrorArea">
        <div id="viewErrorMessageButton" onClick={() => handleClick()}>
          View error message
        </div>
        <ErrorMessageBox />
      </div>
    </>
  );
}

export default ErrorResponse;
