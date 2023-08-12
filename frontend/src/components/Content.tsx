import "./Content.css";
import { useContext } from "react";
import { PageContext, PageContextType } from "../App";

import Download from "./content/Download";
import Gallery from "./content/Gallery";
import Activity from "./content/Activity";
import Settings from "./content/Settings";

function Content() {
  const pageContext: PageContextType = useContext(PageContext);

  let content: JSX.Element;
  if (pageContext.download.downloadState) {
    content = <Download />;
  } else if (pageContext.gallery.galleryState) {
    content = <Gallery />;
  } else if (pageContext.activity.activityState) {
    content = <Activity />;
  } else if (pageContext.settings.settingsyState) {
    content = <Settings />;
  } else {
    content = <Download />;
  }

  return <div id="content">{content}</div>;
}
export default Content;
