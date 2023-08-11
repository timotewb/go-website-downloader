import "./MenuBar.css";
import downloadSVG from "../assets/images/download.svg";
import gallerySVG from "../assets/images/gallery.svg";
import activitySVG from "../assets/images/activity.svg";
import { useState } from "react";

function MenuBar() {
  const [downloadState, setDownloadState] = useState(true);
  const [galleryState, setGalleryState] = useState(false);
  const [activityState, setActivityState] = useState(false);

  const downloadID = "downloadButton";
  const galleryID = "galleryButton";
  const activityID = "activityButton";

  const handleClick = (id: string) => {
    if (id === downloadID) {
      setDownloadState(true);
      setGalleryState(false);
      setActivityState(false);
    } else if (id === galleryID) {
      setDownloadState(false);
      setGalleryState(true);
      setActivityState(false);
    } else if (id === activityID) {
      setDownloadState(false);
      setGalleryState(false);
      setActivityState(true);
    } else {
      setDownloadState(true);
      setGalleryState(false);
      setActivityState(false);
    }
  };

  return (
    <div id="menuBar">
      <div
        id="downloadButton"
        className={downloadState ? "button button-active" : "button"}
        onClick={() => handleClick("downloadButton")}
      >
        <img src={downloadSVG}></img> Download
      </div>
      <div
        id="galleryButton"
        className={galleryState ? "button button-active" : "button"}
        onClick={() => handleClick("galleryButton")}
      >
        <img src={gallerySVG}></img> Gallery
      </div>
      <div
        id="activityButton"
        className={activityState ? "button button-active" : "button"}
        onClick={() => handleClick("activityButton")}
      >
        <img src={activitySVG}></img> Activity
      </div>
    </div>
  );
}
export default MenuBar;
