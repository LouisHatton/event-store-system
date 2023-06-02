import NetworkService from './NetworkService';

export type Connection = {
	id: string;
	urlId: string;
	projectId: string;
	name: string;
	metadata: ConnectionMetadata;
	tags: string[];
	status: string;
	schema: Record<string, string>;
};

export type ConnectionMetadata = {
	createdBy: string;
	createdAt: string;
};

export type NewConnection = {
	name: string;
	tags: string[];
};

export class ConnectionService {
	getConnection(projectId: string, connectionId: string): Promise<Connection> {
		return NetworkService.get(`/api/v1/projects/${projectId}/connections/${connectionId}`);
	}

	listConnections(projectId: string): Promise<Connection[]> {
		return NetworkService.get(`/api/v1/projects/${projectId}/connections`);
	}

	create(projectId: string, connection: NewConnection): Promise<Connection> {
		return NetworkService.post(`/api/v1/projects/${projectId}/connections`, connection);
	}

	set(projectId: string, connectionId: string, connection: NewConnection): Promise<Connection> {
		return NetworkService.post(
			`/api/v1/projects/${projectId}/connections/${connectionId}`,
			connection
		);
	}
}

export default new ConnectionService();
