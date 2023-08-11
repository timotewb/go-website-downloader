import "./MenuBar.css";
import downloadSVG from "../assets/images/download.svg";
import gallerySVG from "../assets/images/gallery.svg";
import activitySVG from "../assets/images/activity.svg";
import asettingsSVG from "../assets/images/settings.svg";
import { useState, useContext } from "react";
import { PageContext, PageContextType } from "../App";

function MenuBar() {
  const pageContext: PageContextType = useContext(PageContext);

  const downloadID = "downloadButton";
  const galleryID = "galleryButton";
  const activityID = "activityButton";
  const settingsID = "settingsButton";

  const handleClick = (id: string) => {
    if (id === downloadID) {
      pageContext.download.setDownloadState(true);
      pageContext.gallery.setGalleryState(false);
      pageContext.activity.setActivityState(false);
      pageContext.settings.setSettingsState(false);
    } else if (id === galleryID) {
      pageContext.download.setDownloadState(false);
      pageContext.gallery.setGalleryState(true);
      pageContext.activity.setActivityState(false);
      pageContext.settings.setSettingsState(false);
    } else if (id === activityID) {
      pageContext.download.setDownloadState(false);
      pageContext.gallery.setGalleryState(false);
      pageContext.activity.setActivityState(true);
      pageContext.settings.setSettingsState(false);
    } else if (id === settingsID) {
      pageContext.download.setDownloadState(false);
      pageContext.gallery.setGalleryState(false);
      pageContext.activity.setActivityState(false);
      pageContext.settings.setSettingsState(true);
    } else {
      pageContext.download.setDownloadState(true);
      pageContext.gallery.setGalleryState(false);
      pageContext.activity.setActivityState(false);
      pageContext.settings.setSettingsState(false);
    }
  };

  return (
    <div id="menuBar">
      <div
        id="downloadButton"
        className={
          pageContext.download.downloadState ? "button button-active" : "button"
        }
        onClick={() => handleClick("downloadButton")}
      >
        <img src={downloadSVG}></img> Download
      </div>
      <div
        id="galleryButton"
        className={
          pageContext.gallery.galleryState ? "button button-active" : "button"
        }
        onClick={() => handleClick("galleryButton")}
      >
        <img src={gallerySVG}></img> Gallery
      </div>
      <div
        id="activityButton"
        className={
          pageContext.activity.activityState ? "button button-active" : "button"
        }
        onClick={() => handleClick("activityButton")}
      >
        <img src={activitySVG}></img> Activity
      </div>
      <div className="spacer"></div>
      <div
        id={settingsID}
        className={
          pageContext.settings.settingsyState
            ? "button button-active"
            : "button"
        }
        onClick={() => handleClick(settingsID)}
      >
        <img src={asettingsSVG}></img> Settings
      </div>
    </div>
  );
}
export default MenuBar;
