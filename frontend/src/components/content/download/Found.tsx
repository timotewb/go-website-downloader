import { lib } from "../../../../wailsjs/go/models";
import "./shared.css";
import successSVG from "../../../assets/images/success.svg";
import { PageContext, PageContextType } from "../../../App";
import { useContext } from "react";
import * as App from "../../../../wailsjs/go/main/App"

function Found(props: lib.ResponseType) {
  const pageContext: PageContextType = useContext(PageContext);

  const handleClick = () => {
    App.GetSite(props).then(() =>{
      pageContext.download.setDownloadState(false);
      pageContext.gallery.setGalleryState(false);
      pageContext.activity.setActivityState(true);
      pageContext.settings.setSettingsState(false);
    });
  };

  return (
    <>
      <div id="messageArea">
        <div className="icon">
          <img src={successSVG}></img>
        </div>
        <div className="message">Found!</div>
      </div>
      <div className="faviconDownloadArea">
        <img id="faviconImg" src={props.favicon_url}></img>
        <div id="downloadNowButton"
          onClick={() => handleClick()}>
          <div id="inputButton">Download</div>
        </div>
      </div>
    </>
  );
}

export default Found;
