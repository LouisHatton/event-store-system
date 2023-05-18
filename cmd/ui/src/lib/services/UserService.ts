import { updateEmail, updateProfile, type User } from 'firebase/auth';
import AuthenticationService from './AuthenticationService';

class UserService {
	async updateUser(email: string, displayName: string) {
		let currentUser = AuthenticationService.currentUser;
		if (!currentUser) throw new Error('Not currently logged in');
		if (currentUser && currentUser?.email !== email) {
			await updateEmail(currentUser, email);
		}
		if (displayName !== currentUser?.displayName) {
			updateProfile(currentUser, { displayName });
		}
	}
}

export default new UserService();
