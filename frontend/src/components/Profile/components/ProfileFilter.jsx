
import ProfilePersonal from "./ProfilePersonalSection";
import ProfileDelivery from "./ProfileDeliverySection";
import ProfilePassword from "./ProfilePasswordSection";
import ProfileLogout from "./ProfileLogout";
import {memo, useState} from "react";
import ProfileApplications from "./ProfileApplications";
import "./style.css"
import {useNavigate} from "react-router";



const ProfileFilter =({profilePage})=>{
   const profile = profilePage
    const navigate = useNavigate();
    const [active, setIsActive] = useState(0)
    const buttons = {
        profile:0,
        items:1,
        role:2,
    }



    const handleActiveButton =(id) =>{
        setIsActive(active === id ? null : id)
    }

    return(
        <>
        <div className = "user-profile-filter-wrapper">
            <div className = "user-profile-filter-controls">
                <button className = {active === buttons.profile ? "active" : ""} onClick={()=>handleActiveButton(buttons.profile)}>Профиль</button>
                <button className = {active === buttons.items ? "active" : ""} onClick={()=>handleActiveButton(buttons.items)}>История заявок</button>
                {profile.userRole === profile.roles.adminUser && (
                <button className = {active === buttons.role ? "active" : ""} onClick={()=>navigate("/admin")}>Админ-панель</button>
                        )}
            </div>
        </div>
    {active === buttons.items &&
    <>
        {profile.user.itemData && active === buttons.items ? (
       <ProfileApplications
       profileApplications={profile}/>
            ):(
                <h1>История заявок отсутствует.</h1>
            )}
    </>
    }
    {active === buttons.profile &&
    <>
        <ProfilePersonal
            profile = {profile}
            setProfile = {profile.setProfile}/>
        <ProfileDelivery
            profile = {profile}
            setProfile = {profile.setProfile}
        />
        <ProfilePassword
            profile = {profile}
            setProfile = {profile.setProfile}/>
        <ProfileLogout />
    </>
    }
        </>
)
}

export default memo(ProfileFilter);