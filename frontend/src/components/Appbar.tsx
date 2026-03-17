import Button from "./SimButton";
import SignupPage from "../pages/Signup";
import LoginPage from "../pages/Login";
import { useState } from "react";

const Appbar = () => {
    const [openSignup,SetOpenSignup] = useState(false);
    const [openLogIn,setOpenLogIn]=useState(false);
  return (
    <>  
        {openSignup && <SignupPage onClose={()=>SetOpenSignup(false)}/>}
        {openLogIn && <LoginPage onClose={()=> setOpenLogIn(false)}/>}
        <div className="h-20 w-full flex justify-between border-gray-200 shadow">
            <div className="flex justify-around p-6 px-11 gap-x-2">
                <div className="text-lg text-red-500 hover:text-red-600 font-bold p-2 ">
                    <i>Wintrest</i>
                </div>
                <div className="text-lg font-bold p-2">
                    <div className=" hover:bg-slate-100 rounded-lg">
                        Explore
                    </div>
                </div>
            </div>
            <div>
                <div className="flex justify-around p-6 px-11">
                    <div className="flex justify-around">
                        <div className="hover:bg-slate-100 rounded-lg text-lg font-bold p-2">
                            About 
                        </div>
                        <div className="hover:bg-slate-100 rounded-lg text-lg font-bold p-2">
                            Businesses
                        </div>
                        <div className="hover:bg-slate-100 rounded-lg text-lg font-bold p-2">
                            Create
                        </div>
                        <div className="text-lg font-bold p-2 hover:bg-slate-100 rounded-lg">
                            News 
                        </div>
                    </div>
                    <div className="flex gap-x-3 ml-1 me-1">
                        <div>
                            <Button onClick={()=>{
                                setOpenLogIn(true);
                            }} text="Log In " colour="bg-red-600" textColour="text-white"></Button>
                        </div>
                        <div>
                            <Button onClick={()=>{
                                SetOpenSignup(true)
                            }} text="Sign Up" colour="bg-slate-200" textColour="text-black"></Button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </>
  );
};

export default Appbar;  