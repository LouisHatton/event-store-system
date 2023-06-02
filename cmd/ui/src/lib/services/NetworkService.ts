import AuthenticationService from './AuthenticationService';

export type ApiError = {
	status: number;
	message: string;
};

function handleNetworkError(r: Response): ApiError {
	return {
		status: r.status,
		message: 'unable to parse error'
	};
}

class NetworkService {
	async fetch<T>(route: string, method = 'GET', body?: object) {
		let token = await AuthenticationService.getToken();
		if (!token) throw new Error('not logged in');

		let response = await fetch(route, {
			method,
			body: body ? JSON.stringify(body) : undefined,
			headers: {
				Authorization: 'Bearer ' + token,
				'content-type': 'application/json'
			}
		});

		if (!response.ok) {
			throw handleNetworkError(response);
		}

		return (await response.json()) as T;
	}

	async get<T>(route: string) {
		return this.fetch(route, 'GET') as T;
	}

	async post<T>(route: string, body: object) {
		return this.fetch(route, 'POST', body) as T;
	}
}

export default new NetworkService();
