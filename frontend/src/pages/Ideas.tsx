import { useState, type JSX } from "react";
import Sidebar from "../components/SideBar";
import Appbar from "../components/Appbar";
import Header from "../components/Header";
import Subheader from "../components/Subheading";
import LoginPage from "./Login";
import SignupPage from "./Signup";

// global state 
import { useAuthStore } from "../store/atoms/AuthStat";

const Ideas = ():JSX.Element=>{
    const [loginOpen,setLoginOpen]=useState(false);
    const [signupOpen,setSignupOpen]=useState(false);
    const LoggedIn = useAuthStore((s)=>s.isLoggedIn);

    return(<>
    <div className="flex h-screen">
        {!LoggedIn && <LoginPage onClose={()=>{}} openSignup={()=>{
            setLoginOpen(false);
            setSignupOpen(true);
        }}/>}
        {loginOpen && <LoginPage onClose={()=>setLoginOpen(false)} openSignup={()=>{
            setLoginOpen(false);
            setSignupOpen(true);
        }}/>}
        {signupOpen && <SignupPage onClose={()=>{
            setSignupOpen(false);
        }} openLogin={()=>{
            setSignupOpen(false);
            setLoginOpen(true);
        }}/>}
        
            <Sidebar/>
        <div className=" relative flex flex-col items-center w-full gap-8">
            <Appbar/>
            <Header text="Explore Your Ideas!"/>
            <Subheader text="Search the images based on your interests and hobbies :) "/>
        </div>
    </div>
    </>)
}

export default Ideas;