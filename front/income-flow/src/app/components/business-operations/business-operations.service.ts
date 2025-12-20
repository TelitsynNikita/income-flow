import {inject, Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';

interface Contractor {
  name: string;
  inn: string;
  type: string;
}

export interface BusinessOperation {
  id: number;
  date: string;
  operation_type: string;
  operation_id: number;
  goods_count: number;
  goods_id: number;
  section_id: number;
  contractors_id: number;
  contract: Contractor;
}

@Injectable({providedIn: 'root'})
export class BusinessOperationsService {
  private http = inject(HttpClient);
  private url = 'http://localhost:8080'


  getOperations(): Observable<BusinessOperation[]> {
    return this.http.get<BusinessOperation[]>(`${this.url}/operations/get_business_operations`);
  }
}
