import AdminPageHeader from "../components/AdminPage/components/AdminPageHeader";
import {GetAdminPage} from "../components/AdminPage/hooks/GetAdminPage";
import AdminPageFilter from "../components/AdminPage/components/AdminPageFilter";
import "../components/AdminPage/components/style.css"
import {useAuth} from "../context/AuthContext";
import PageNotFound from "../features/404NotFound/PageNotFound";
import ForbiddenPage from "../features/403Forbidden/ForbiddenPage";




const AdminPage = () =>{
    const adminInfo = GetAdminPage()

    const user = useAuth()

    if (!user.user){
        return <ForbiddenPage />
    }



    if (adminInfo.loading){
        return <div>loading...</div>
    }

    if (adminInfo.error){
        return <PageNotFound />
    }
    return (
        <div className="admin_page-wrapper">
            <div className="admin_page">
                <AdminPageHeader
                    adminInfo = {adminInfo}
                />
                <AdminPageFilter/>
            </div>
        </div>
    )
}

export default AdminPage;