import { useState, useEffect } from "react";
import UserService from "../services/UserService";
import User from "../models/User";

export default function ListUserComponent(): JSX.Element {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    UserService.getUsers().then((fetchedUsers: User[]) => {
      setUsers(fetchedUsers);
    });
  }, []);

  return (
    <div>
      <h2 className="text-center">Users</h2>
      <div className="row">
        <table className="table table-striped table-bordered">
          <thead>
            <tr>
              <th>Name</th>
              <th>Gender</th>
              <th>Email</th>
              <th>Password</th>
              <th>Phone Number</th>
              <th>Date of Birth</th>
              <th>Role</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <tr key={user.userId}>
                <td>{user.name}</td>
                <td>{user.gender}</td>
                <td>{user.email}</td>
                <td>{user.password}</td>
                <td>{user.phoneNumber}</td>
                <td>{user.dateOfBirth}</td>
                <td>{user.role}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
