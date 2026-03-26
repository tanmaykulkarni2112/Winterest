import { create } from "zustand";

type catagoryProp = {
    catagoryName:string;
    setCatagoryName:(catagory:string)=>void;
}

export const UseCatagoryProp = create<catagoryProp>((set)=>({
    catagoryName:"",
    setCatagoryName:((catagory)=>set({catagoryName:catagory}))
}))