import { useState, useEffect, useRef } from "react";
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
  const inputRef = useRef("");

  const check = () => {
    const handleClick = (site: string) => {
      setSiteName(site);
      pageContext.gallery.setShowSiteList(true);
    };

    App.ListGallery().then((data) => {
      console.log("App.ListGallery()");

      data.sort((a: models.GalleryType, b: models.GalleryType) =>a.site_name.localeCompare(b.site_name));
      const gallery = data.map((s: models.GalleryType, i: number) => {
        return (
          <div className="galleryCell" key={i}>
            <span
              onClick={() => handleClick(s.site_name)}
              className="gallerySiteSelect"
            >
              <img id="faviconImg" src={s.favicon}></img>
              <br />
              {s.site_name}
            </span>
          </div>
        );
      });
      setGallery(gallery);
    });
  };
  const renderContent = () => {
    if (pageContext.gallery.showSiteList) {
      return <ListSite site_name={siteName} />;
    }
    const handleBlur = (e: React.FocusEvent<HTMLInputElement, Element>) => {
      e.target.placeholder = "Enter URL";
    };
    const handleFocus = (e: React.FocusEvent<HTMLInputElement, Element>) => {
      e.target.placeholder = "";
    };
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      inputRef.current = e.target.value;
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
        </div>
        <div id="gallery">{gallery}</div>
      </>
    );
  };

  useEffect(() => {
    check();
    renderContent();
  }, []);


  return renderContent();
}

export default Gallery;
