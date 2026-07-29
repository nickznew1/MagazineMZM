


const AdminPageFilterUsers = ({users, loading}) =>{

    return (
        <div>
        <table className="admin_page-users-table">
            <thead>
            <tr>
                <th>ID</th>
                <th>Логин</th>
                <th>Роль</th>
                <th>Почта</th>
                <th>Дата регистрации</th>
            </tr>
            </thead>
        </table>
            <div className = "admin_page-table-body">
            <table className="admin_page-users-table">
            <tbody>
            {users?.map((user) => (
                <tr key={user.id}>
                    <td>{user.id}</td>
                    <td>{user.login}</td>
                    <td>{user.user_role}</td>
                    <td>{user.email}</td>
                    <td>
                        {new Date(user.registration_date)
                            .toLocaleDateString("ru-RU")}
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
        </div>
        </div>
    )
}

export default AdminPageFilterUsers;