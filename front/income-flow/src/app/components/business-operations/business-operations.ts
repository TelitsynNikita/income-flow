import {Component, inject} from '@angular/core';
import {BusinessOperation, BusinessOperationsService} from './business-operations.service';
import {App} from '../../app';

@Component({
  selector: 'app-business-operations',
  imports: [],
  templateUrl: './business-operations.html',
  styleUrl: './business-operations.scss',
})
export class BusinessOperations {
  readonly businessOperationForm = inject(BusinessOperationsService);
  operations: BusinessOperation[] = [];

  constructor(private app: App) {
    this.businessOperationForm.getOperations()
      .subscribe(res => {
        this.operations = res;
      })
  }
}
