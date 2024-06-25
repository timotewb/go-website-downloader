import {useState, useEffect} from "react";
import "./Gallery.css";
import "./shared.css";
import * as App from "../../../wailsjs/go/main/App";
import { models } from "../../../wailsjs/go/models";

function Gallery() {
  const [gallery, setGallery] = useState<JSX.Element[] | null>(null);

  const check = () => {
  App.ListGallery().then((data)=>{
    console.log("App.ListGallery()");
    const gallery = data.map((s: models.GalleryType, i: number) => {
      return (
        <div className="galleryCell" key={i}>
          <img id="faviconImg" src={s.favicon}></img>
          {s.site_name}
        </div>
      )
    });
    setGallery(gallery);
  });
};
useEffect(() => {check();}, []);

  return(
    <div id="gallery">
      {gallery}
    </div>
  );
}

export default Gallery;
