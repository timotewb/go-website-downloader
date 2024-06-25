// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';
import {lib} from '../models';

export function CheckActivity():Promise<models.CheckActivityType>;

export function FindURL(arg1:string):Promise<lib.ResponseType>;

export function GetSettings():Promise<models.SettingsType>;

export function GetSite(arg1:models.ResponseType):Promise<void>;

export function ListGallery():Promise<models.ListGalleryType>;

export function RemoveStaleActivity(arg1:string,arg2:string):Promise<void>;

export function ShutdownContentDirWebServer():Promise<void>;
