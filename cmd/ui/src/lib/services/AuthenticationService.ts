import { env } from '$env/dynamic/public';
import { authStore } from '$lib/stores/authStore';
import { initializeApp, type FirebaseApp, type FirebaseOptions } from 'firebase/app';
import {
	getAuth,
	type Auth,
	signInWithEmailAndPassword,
	type UserCredential,
	type User,
	createUserWithEmailAndPassword,
	GoogleAuthProvider,
	signInWithPopup,
	signInWithRedirect
} from 'firebase/auth';

const firebaseConfig: FirebaseOptions = {
	apiKey: env['PUBLIC_API_KEY'],
	authDomain: env['PUBLIC_AUTH_DOMAIN']
};

class AuthenticationService {
	app: FirebaseApp;
	auth: Auth;
	googleProvider: GoogleAuthProvider;
	public currentUser: User | null = null;

	constructor() {
		this.app = initializeApp(firebaseConfig);
		this.auth = getAuth(this.app);
		this.googleProvider = new GoogleAuthProvider();

		this.auth.onAuthStateChanged((user) => {
			this.currentUser = user;
			authStore.set({
				isLoggedIn: user != null,
				user,
				initialised: true
			});
		});
	}

	async signInWithPassword(email: string, password: string): Promise<UserCredential> {
		return signInWithEmailAndPassword(this.auth, email, password);
	}

	async registerWithPassword(email: string, password: string): Promise<UserCredential> {
		return createUserWithEmailAndPassword(this.auth, email, password);
	}

	signInWithGoogle() {
		return signInWithRedirect(this.auth, this.googleProvider);
	}

	async isLoggedIn() {
		return this.auth.currentUser != null;
	}

	async logOut() {
		return this.auth.signOut();
	}

	async getToken() {
		return this.auth.currentUser?.getIdToken();
	}
}

export default new AuthenticationService();
