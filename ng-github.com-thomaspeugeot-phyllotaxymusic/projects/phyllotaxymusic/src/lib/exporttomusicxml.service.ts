// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { ExportToMusicxmlAPI } from './exporttomusicxml-api'
import { ExportToMusicxml, CopyExportToMusicxmlToExportToMusicxmlAPI } from './exporttomusicxml'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { ParameterAPI } from './parameter-api'

@Injectable({
  providedIn: 'root'
})
export class ExportToMusicxmlService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ExportToMusicxmlServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private exporttomusicxmlsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.exporttomusicxmlsUrl = origin + '/api/github.com/thomaspeugeot/phyllotaxymusic/go/v1/exporttomusicxmls';
  }

  /** GET exporttomusicxmls from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<ExportToMusicxmlAPI[]> {
    return this.getExportToMusicxmls(GONG__StackPath, frontRepo)
  }
  getExportToMusicxmls(GONG__StackPath: string, frontRepo: FrontRepo): Observable<ExportToMusicxmlAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<ExportToMusicxmlAPI[]>(this.exporttomusicxmlsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<ExportToMusicxmlAPI[]>('getExportToMusicxmls', []))
      );
  }

  /** GET exporttomusicxml by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ExportToMusicxmlAPI> {
    return this.getExportToMusicxml(id, GONG__StackPath, frontRepo)
  }
  getExportToMusicxml(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ExportToMusicxmlAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.exporttomusicxmlsUrl}/${id}`;
    return this.http.get<ExportToMusicxmlAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched exporttomusicxml id=${id}`)),
      catchError(this.handleError<ExportToMusicxmlAPI>(`getExportToMusicxml id=${id}`))
    );
  }

  // postFront copy exporttomusicxml to a version with encoded pointers and post to the back
  postFront(exporttomusicxml: ExportToMusicxml, GONG__StackPath: string): Observable<ExportToMusicxmlAPI> {
    let exporttomusicxmlAPI = new ExportToMusicxmlAPI
    CopyExportToMusicxmlToExportToMusicxmlAPI(exporttomusicxml, exporttomusicxmlAPI)
    const id = typeof exporttomusicxmlAPI === 'number' ? exporttomusicxmlAPI : exporttomusicxmlAPI.ID
    const url = `${this.exporttomusicxmlsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ExportToMusicxmlAPI>(url, exporttomusicxmlAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<ExportToMusicxmlAPI>('postExportToMusicxml'))
    );
  }
  
  /** POST: add a new exporttomusicxml to the server */
  post(exporttomusicxmldb: ExportToMusicxmlAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ExportToMusicxmlAPI> {
    return this.postExportToMusicxml(exporttomusicxmldb, GONG__StackPath, frontRepo)
  }
  postExportToMusicxml(exporttomusicxmldb: ExportToMusicxmlAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ExportToMusicxmlAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ExportToMusicxmlAPI>(this.exporttomusicxmlsUrl, exporttomusicxmldb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted exporttomusicxmldb id=${exporttomusicxmldb.ID}`)
      }),
      catchError(this.handleError<ExportToMusicxmlAPI>('postExportToMusicxml'))
    );
  }

  /** DELETE: delete the exporttomusicxmldb from the server */
  delete(exporttomusicxmldb: ExportToMusicxmlAPI | number, GONG__StackPath: string): Observable<ExportToMusicxmlAPI> {
    return this.deleteExportToMusicxml(exporttomusicxmldb, GONG__StackPath)
  }
  deleteExportToMusicxml(exporttomusicxmldb: ExportToMusicxmlAPI | number, GONG__StackPath: string): Observable<ExportToMusicxmlAPI> {
    const id = typeof exporttomusicxmldb === 'number' ? exporttomusicxmldb : exporttomusicxmldb.ID;
    const url = `${this.exporttomusicxmlsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<ExportToMusicxmlAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted exporttomusicxmldb id=${id}`)),
      catchError(this.handleError<ExportToMusicxmlAPI>('deleteExportToMusicxml'))
    );
  }

  // updateFront copy exporttomusicxml to a version with encoded pointers and update to the back
  updateFront(exporttomusicxml: ExportToMusicxml, GONG__StackPath: string): Observable<ExportToMusicxmlAPI> {
    let exporttomusicxmlAPI = new ExportToMusicxmlAPI
    CopyExportToMusicxmlToExportToMusicxmlAPI(exporttomusicxml, exporttomusicxmlAPI)
    const id = typeof exporttomusicxmlAPI === 'number' ? exporttomusicxmlAPI : exporttomusicxmlAPI.ID
    const url = `${this.exporttomusicxmlsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<ExportToMusicxmlAPI>(url, exporttomusicxmlAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<ExportToMusicxmlAPI>('updateExportToMusicxml'))
    );
  }

  /** PUT: update the exporttomusicxmldb on the server */
  update(exporttomusicxmldb: ExportToMusicxmlAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ExportToMusicxmlAPI> {
    return this.updateExportToMusicxml(exporttomusicxmldb, GONG__StackPath, frontRepo)
  }
  updateExportToMusicxml(exporttomusicxmldb: ExportToMusicxmlAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ExportToMusicxmlAPI> {
    const id = typeof exporttomusicxmldb === 'number' ? exporttomusicxmldb : exporttomusicxmldb.ID;
    const url = `${this.exporttomusicxmlsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<ExportToMusicxmlAPI>(url, exporttomusicxmldb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated exporttomusicxmldb id=${exporttomusicxmldb.ID}`)
      }),
      catchError(this.handleError<ExportToMusicxmlAPI>('updateExportToMusicxml'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in ExportToMusicxmlService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("ExportToMusicxmlService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}