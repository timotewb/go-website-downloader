import { lib } from "../../../../wailsjs/go/models";
import "./shared.css";
import successSVG from "../../../assets/images/success.svg";
import Cookies from 'universal-cookie';
import { PageContext, PageContextType } from "../../../App";
import { useContext } from "react";

function Found(props: lib.ResponseType) {
  const cookies = new Cookies();
  const pageContext: PageContextType = useContext(PageContext);

  const handleClick = () => {
      pageContext.download.setDownloadState(false);
      pageContext.gallery.setGalleryState(false);
      pageContext.activity.setActivityState(true);
      pageContext.settings.setSettingsState(false);
      cookies.set('myCat1', 'Pacman', { path: '/' });
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
