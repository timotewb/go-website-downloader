import {useState, useEffect} from "react";
import "./Gallery.css";
import "./shared.css";
import * as App from "../../../wailsjs/go/main/App";
import { models } from "../../../wailsjs/go/models";
import { PageContext, PageContextType } from "../../App";
import { useContext } from "react";
import ListSite from "./gallery/ListSite";

function Gallery() {
  const pageContext: PageContextType = useContext(PageContext);
  const [gallery, setGallery] = useState<JSX.Element[] | null>(null);
  const [siteName, setSiteName] = useState("");

  const check = () => {

  const handleClick = (site: string) => {
    setSiteName(site);
    pageContext.gallery.setShowSiteList(true);
  };

  App.ListGallery().then((data)=>{
    console.log("App.ListGallery()");
    const gallery = data.map((s: models.GalleryType, i: number) => {
      return (
        <div className="galleryCell" key={i}>
            <span onClick={() => handleClick(s.site_name)} className="gallerySiteSelect">
              <img id="faviconImg" src={s.favicon}></img>
              <br />
              {s.site_name}
            </span>
        </div>
      )
    });
    setGallery(gallery);
  });
};
useEffect(() => {check();}, []);

  const renderContent = () => {
    if (pageContext.gallery.showSiteList){
      return <ListSite site_name={siteName} />;
    }
    return gallery;
  };
  return(
    <div id="gallery">
      {renderContent()}
    </div>
  );
}

export default Gallery;
