

const AdminPageHeader = ({adminInfo}) => {
    const adminPage = adminInfo;
    return(
        <div className="admin_page-header">
            <h1>Админ панель пользователя : {adminPage.user.login}</h1>
        </div>
    )
}

export default AdminPageHeader;