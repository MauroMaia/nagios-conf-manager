import {HttpClient} from '@angular/common/http';
import {Injectable, OnInit} from '@angular/core';
import {Observable} from 'rxjs';
import {Command} from '../../model/command';
import {GenericDataService} from '../GenericDataService';

@Injectable({
  providedIn: 'root',
})
export class CommandsService extends GenericDataService {
  private readonly getAllCommandsUrl = this.server + '/commands';

  constructor(private http: HttpClient) {
    super();
  }

  getAllCommands(): Observable<Command[]> {
    return this.http.get<Command[]>(this.getAllCommandsUrl);
    /* .pipe(
       tap(_ => console.info('fetched hosts')),
       catchError(this.handleError<Host[]>('getHeroes', []))
     );*/
  }
}
