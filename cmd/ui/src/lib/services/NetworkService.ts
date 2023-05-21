import AuthenticationService from './AuthenticationService';

class NetworkService {
	async get<T>(route: string, method = 'GET') {
		let token = await AuthenticationService.getToken();

		let response = await fetch(route, {
			method,
			headers: {
				Authorization: 'Bearer ' + token
			}
		});

		return (await response.json()) as T;
	}
}

export default new NetworkService();
