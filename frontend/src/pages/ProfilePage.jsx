import {memo} from "react"
import {UserProfilePage} from "../components/Profile/hooks/userProfilePage";
import "../components/Profile/components/style.css";
import ProfileFilter from "../components/Profile/components/ProfileFilter";
import NotAuthorized from "../features/NotAuthorized/NotAuthorized";


const ProfilePage = () => {
    const profile = UserProfilePage()

    if (profile.loading) {
        return <div>loading...</div>

    }
    if (profile.error) {
        return <NotAuthorized/>
    }

  return (
      <div className = "user-profile-wrapper">
          <div className ="user-profile">
              <ProfileFilter profilePage={profile}/>
      </div>
      </div>
  );

}

export default memo(ProfilePage)