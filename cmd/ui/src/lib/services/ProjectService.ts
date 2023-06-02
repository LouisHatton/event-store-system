import NetworkService from './NetworkService';

export type Project = {
	id: string;
	name: string;
	metadata: ProjectMetadata;
	config: ProjectConfig;
	users: string[];
	adminUsers: string[];
	roUsers: string[];
	rwUsers: string[];
};

export type ProjectConfig = {
	colour: string;
};

export type ProjectMetadata = {
	createdBy: string;
	createdAt: string;
};

export class ProjectService {
	listProjects(): Promise<Project[]> {
		return NetworkService.get('/api/v1/projects');
	}
}

export default new ProjectService();
