import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable, of} from 'rxjs';
import {map} from 'rxjs/operators';
import {TimePeriod} from '../../model/time-periods';

@Injectable({
  providedIn: 'root',
})
export class TimePeriodService {

  server = 'http://127.0.0.1:3000/timeperiods';

  constructor(private http: HttpClient) {
  }

  getAllHosts(): Observable<TimePeriod[]> {
    return this.http.get<any[]>(this.server).pipe<TimePeriod[]>(
    // @ts-ignore
      map<TimePeriod[]>(x => {
        return x.map(object => {
          const tp: TimePeriod = {
            name           : object.name,
            timeperiod_name: object.timeperiod_name,
            alias          : object.alias,
            sunday         : object.sunday,
            monday         : object.monday,
            tuesday        : object.tuesday,
            wednesday      : object.wednesday,
            thursday       : object.thursday,
            friday         : object.friday,
            saturday       : object.saturday,
          };
          return tp;
        });
      })
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      console.error(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}



